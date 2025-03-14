// profile-converter.ts
import { Profile, Location, Function, Line, Sample, } from '@/utils/pprof';

export interface FlameNode {
  name: string;
  value: number;
  children: FlameNode[];
  // 可扩展的元数据字段
  meta?: {
    file?: string;
    line?: number;
    systemName?: string;
    module?: string;
  };
}

export function profileToFlameGraph(profile: Profile): FlameNode {
  // 1. 初始化根节点
  const root: FlameNode = {
    name: "root",
    value: 0,
    children: [],
  };

  // 2. 遍历所有样本
  for (const sample of profile.Sample) {
    // 3. 获取调用栈（从深到浅）
    const stack = sample.Location
      .flatMap(loc => getLocationStack(loc))
      .reverse(); // 反转获取从根到叶的顺序

    // 4. 累加样本值到调用树
    let currentNode = root;
    currentNode.value += getSampleValue(sample);

    for (const frame of stack) {
      const childName = getFrameName(frame);
      let child = currentNode.children.find(c => c.name === childName);

      if (!child) {
        child = {
          name: childName,
          value: 0,
          children: [],
          meta: getFrameMeta(frame)
        };
        currentNode.children.push(child);
      }

      child.value += getSampleValue(sample);
      currentNode = child;
    }
  }

  return root;
}

// 辅助函数：获取样本的主数值
function getSampleValue(sample: Sample): number {
  // 使用默认样本类型或第一个值
  return sample.Value[0] || 0; // 根据实际样本类型调整
}

// 辅助函数：展开Location的完整调用栈
function getLocationStack(location: Location): Array<{ func?: Function; line?: Line }> {
    const stack = [];
    const currentLocation: Location | null = location;
  
    while (currentLocation) {
      if (currentLocation.IsFolded) {
        // 生成折叠节点，避免使用new Function可能的问题
        const foldedFunc: Function = {
          ID: 0,
          Name: "[folded]",
          SystemName: "",
          Filename: "",
          StartLine: 0,
        };
        // 添加当前Line（如果存在）
        const line = currentLocation.Line.length > 0 ? currentLocation.Line[0] : undefined;
        stack.push({ func: foldedFunc, line });
        break; // 折叠后停止处理当前Location的后续Line或父节点
      } else {
        // 处理所有Line条目（内联调用）
        for (const line of currentLocation.Line) {
          stack.push({
            func: line.Function || undefined,
            line: line
          });
        }
        // 假设当前数据结构无法获取父Location，停止处理
        break;
      }
    }
  
    return stack;
  }
// 辅助函数：生成节点显示名称
function getFrameName(frame: {
    func?: Function;
    line?: Line;
  }): string {
    if (!frame.func) return "unknown";
    
    // 折叠节点特殊处理
    if (frame.func.Name === "[folded]") {
      return "[folded]";
    }
  
    const baseName = frame.func.Name || frame.func.SystemName;
    const lineInfo = frame.line 
      ? ` (${frame.func.Filename}:${frame.line.Line})` 
      : "";
    
    return `${baseName}${lineInfo}`;
  }

// 辅助函数：生成元数据
function getFrameMeta(frame: {
  func?: Function;
  line?: Line;
}): FlameNode['meta'] {
  if (!frame.func) return {};

  return {
    file: frame.func.Filename,
    line: frame.line?.Line,
    systemName: frame.func.SystemName,
    module: frame.func.Filename?.split('/').pop()
  };
}
