// profile.ts

export class Profile {
  SampleType: ValueType[] = [];
  DefaultSampleType: string = "";
  Sample: Sample[] = [];
  Mapping: Mapping[] = [];
  Location: Location[] = [];
  Function: Function[] = [];
  Comments: string[] = [];
  DocURL: string = "";
  DropFrames: string = "";
  KeepFrames: string = "";
  TimeNanos: number = 0;
  DurationNanos: number = 0;
  PeriodType: ValueType | null = null;
  Period: number = 0;

  private encodeMu: { lock: () => void; unlock: () => void } = {
    lock: () => {},
    unlock: () => {},
  };

  private commentX: number[] = [];
  private docURLX: number = 0;
  private dropFramesX: number = 0;
  private keepFramesX: number = 0;
  private stringTable: string[] = [];
  private defaultSampleTypeX: number = 0;

  // 解析函数
  constructor(data?: Partial<Profile>) {
    if (data) {
        this.SampleType = (data.SampleType || []).map(v => new ValueType(v));
        this.DefaultSampleType = data.DefaultSampleType || "";
        this.Sample = (data.Sample || []).map(s => new Sample(s));
        this.Mapping = (data.Mapping || []).map(m => new Mapping(m));
        this.Location = (data.Location || []).map(l => new Location(l));
        this.Function = (data.Function || []).map(f => new Function(f));
        this.Comments = data.Comments || [];
        this.DocURL = data.DocURL || "";
        this.DropFrames = data.DropFrames || "";
        this.KeepFrames = data.KeepFrames || "";
        this.TimeNanos = data.TimeNanos || 0;
        this.DurationNanos = data.DurationNanos || 0;
        this.PeriodType = data.PeriodType ? new ValueType(data.PeriodType) : null;
        this.Period = data.Period || 0;
    }
}
  getSampleHeader(): string[] {
    let sh1 = "";
    this.SampleType.forEach((s) => {
        const dflt = s.Type === this.DefaultSampleType ? "[dflt]" : "";
        sh1 += `${s.Type}(${s.Unit})${dflt} `;
      });
    return sh1.trim().split(" ");
  }
  getSampleItems(): Record<string,string>[] {
    const ss: Record<string,string>[]  = [];
    this.Sample.forEach((s) => {
      ss.push(s.getItems());
    });
    return ss;
  }
  getLocationItems(): Record<string,string>[] {
    const ss: Record<string,string>[]  = [];
    this.Location.forEach((l) => {
      ss.push(l.getItems());
    });
    return ss;
  }
  getMappingItems(): Record<string,string>[] {
    const ss: Record<string,string>[]  = [];
    this.Mapping.forEach((m) => {
      ss.push(m.getItems());
    });
    return ss;
  }
  getTitile(): string {
    const ss: string[] = [];
    if (this.Comments){
      this.Comments.forEach((c) => ss.push(`Comment: ${c}`));
    }
    if (this.DocURL) {
      ss.push(`Doc: ${this.DocURL}`);
    }

    if (this.PeriodType) {
      ss.push(`PeriodType: ${this.PeriodType.Type} ${this.PeriodType.Unit}`);
    }

    ss.push(`Period: ${this.Period}`);

    if (this.TimeNanos !== 0) {
      ss.push(`Time: ${new Date(this.TimeNanos / 1e6).toISOString()}`);
    }

    if (this.DurationNanos !== 0) {
      ss.push(`Duration: ${this.DurationNanos / 1e9}s`);
    }
    return ss.join("\n");
  }
  toString(): string {
    const ss: string[] = [];
    ss.push(this.getTitile());

    ss.push("Samples:");

    let sh1 = "";
    this.SampleType.forEach((s) => {
      const dflt = s.Type === this.DefaultSampleType ? "[dflt]" : "";
      sh1 += `${s.Type}/${s.Unit}${dflt} `;
    });
    ss.push(sh1.trim());

    this.Sample.forEach((s) => ss.push(s.toString()));

    ss.push("Locations");
    this.Location.forEach((l) => ss.push(l.toString()));

    ss.push("Mappings");
    this.Mapping.forEach((m) => ss.push(m.toString()));

    return ss.join("\n") + "\n";
  }

  setLabel(key: string, value: string[]): void {
    this.Sample.forEach((sample) => {
      if (!sample.Label) {
        sample.Label = {};
      }
      sample.Label[key] = value;
    });
  }

  removeLabel(key: string): void {
    this.Sample.forEach((sample) => {
      if (sample.Label) {
        delete sample.Label[key];
      }
    });
  }

  checkValid(): Error | null {
    // Check that sample values are consistent
    const sampleLen = this.SampleType.length;
    if (sampleLen === 0 && this.Sample.length !== 0) {
      return new Error("missing sample type information");
    }
    for (const s of this.Sample) {
      if (!s) {
        return new Error("profile has nil sample");
      }
      if (s.Value.length !== sampleLen) {
        return new Error(
          `mismatch: sample has ${s.Value.length} values vs. ${sampleLen} types`
        );
      }
      for (const l of s.Location) {
        if (!l) {
          return new Error("sample has nil location");
        }
      }
    }

    // Check that all mappings/locations/functions are in the tables
    const mappings: { [key: number]: Mapping } = {};
    for (const m of this.Mapping) {
      if (!m) {
        return new Error("profile has nil mapping");
      }
      if (m.ID === 0) {
        return new Error("found mapping with reserved ID=0");
      }
      if (mappings[m.ID]) {
        return new Error(`multiple mappings with same id: ${m.ID}`);
      }
      mappings[m.ID] = m;
    }

    const functions: { [key: number]: Function } = {};
    for (const f of this.Function) {
      if (!f) {
        return new Error("profile has nil function");
      }
      if (f.ID === 0) {
        return new Error("found function with reserved ID=0");
      }
      if (functions[f.ID]) {
        return new Error(`multiple functions with same id: ${f.ID}`);
      }
      functions[f.ID] = f;
    }

    const locations: { [key: number]: Location } = {};
    for (const l of this.Location) {
      if (!l) {
        return new Error("profile has nil location");
      }
      if (l.ID === 0) {
        return new Error("found location with reserved id=0");
      }
      if (locations[l.ID]) {
        return new Error(`multiple locations with same id: ${l.ID}`);
      }
      locations[l.ID] = l;
      if (l.Mapping) {
        if (l.Mapping.ID === 0 || mappings[l.Mapping.ID] !== l.Mapping) {
          return new Error(
            `inconsistent mapping ${l.Mapping}: ${l.Mapping.ID}`
          );
        }
      }
      for (const ln of l.Line) {
        if (!ln.Function) {
          return new Error(
            `location id: ${l.ID} has a line with nil function`
          );
        }
        if (ln.Function.ID === 0 || functions[ln.Function.ID] !== ln.Function) {
          return new Error(
            `inconsistent function ${ln.Function}: ${ln.Function.ID}`
          );
        }
      }
    }

    return null;
  }
}

export class ValueType {
  Type: string = "";
  Unit: string = "";

  private typeX: number = 0;
  private unitX: number = 0;
  constructor(data?: Partial<ValueType>) {
    if (data) {
        this.Type = data.Type || "";
        this.Unit = data.Unit || "";
    }
}
}

export class Sample {
  Location: Location[] = [];
  Value: number[] = [];
  Label: { [key: string]: string[] } = {};
  NumLabel: { [key: string]: number[] } = {};
  NumUnit: { [key: string]: string[] } = {};

  private locationIDX: number[] = [];
  private labelX: { keyX: number; strX: number; numX: number; unitX: number }[] =
    [];

    constructor(data?: Partial<Sample>) {
        if (data) {
            this.Location = (data.Location || []).map(l => new Location(l));
            this.Value = data.Value || [];
            this.Label = data.Label || {};
            this.NumLabel = data.NumLabel || {};
            this.NumUnit = data.NumUnit || {};
        }
    }
  getItems(): { [key: string]: string } {
    const ss: string[] = [];

    this.Value.forEach((v) => {
      ss.push(`${v.toString().padStart(10)}`);
    });
    ss[ss.length-1] += ": ";
    this.Location.forEach((l) => {
      ss[ss.length-1] += `${l.ID} `;
    });
    if (this.Label && Object.keys(this.Label).length > 0) {
      ss.push(this.labelsToString(this.Label));
    }
    if (Object.keys(this.NumLabel).length > 0) {
      ss.push(
        this.numLabelsToString(this.NumLabel, this.NumUnit)
      );
    }
    return ss.reduce((acc, cur, index) => {
      acc[`${index}`] = cur;
      return acc;
    }, {} as { [key: string]: string })
  }
  toString(): string {
    const ss: string[] = [];
    let sv = "";

    this.Value.forEach((v) => {
      sv += ` ${v.toString().padStart(10)}`;
    });
    sv += ": ";
    this.Location.forEach((l) => {
      sv += `${l.ID} `;
    });
    ss.push(sv.trim());

    const labelHeader = "                ";
    if (this.Label && Object.keys(this.Label).length > 0) {
      ss.push(labelHeader + this.labelsToString(this.Label));
    }
    if (Object.keys(this.NumLabel).length > 0) {
      ss.push(
        labelHeader + this.numLabelsToString(this.NumLabel, this.NumUnit)
      );
    }
    return ss.join("\n");
  }

  hasLabel(key: string, value: string): boolean {
    return this.Label[key]?.includes(value) || false;
  }

  private labelsToString(labels: { [key: string]: string[] }): string {
    const ls: string[] = [];
    Object.keys(labels).forEach((k) => {
      ls.push(`${k}:${labels[k].join(",")}`);
    });
    ls.sort();
    return ls.join(" ");
  }

  private numLabelsToString(
    numLabels: { [key: string]: number[] },
    numUnits: { [key: string]: string[] }
  ): string {
    const ls: string[] = [];
    Object.keys(numLabels).forEach((k) => {
      const units = numUnits[k];
      let labelString: string;

      if (units && units.length === numLabels[k].length) {
        const values = numLabels[k].map((v, i) => `${v} ${units[i]}`);
        labelString = `${k}:${values.join(",")}`;
      } else {
        labelString = `${k}:${numLabels[k].join(",")}`;
      }

      ls.push(labelString);
    });
    ls.sort();
    return ls.join(" ");
  }
}

export class Mapping {
  ID: number = 0;
  Start: number = 0;
  Limit: number = 0;
  Offset: number = 0;
  File: string = "";
  BuildID: string = "";
  HasFunctions: boolean = false;
  HasFilenames: boolean = false;
  HasLineNumbers: boolean = false;
  HasInlineFrames: boolean = false;
  KernelRelocationSymbol: string = "";

  private fileX: number = 0;
  private buildIDX: number = 0;

  constructor(data?: Partial<Mapping>) {
        if (data) {
            this.ID = data.ID || 0;
            this.Start = data.Start || 0;
            this.Limit = data.Limit || 0;
            this.Offset = data.Offset || 0;
            this.File = data.File || "";
            this.BuildID = data.BuildID || "";
            this.HasFunctions = data.HasFunctions || false;
            this.HasFilenames = data.HasFilenames || false;
            this.HasLineNumbers = data.HasLineNumbers || false;
            this.HasInlineFrames = data.HasInlineFrames || false;
            this.KernelRelocationSymbol = data.KernelRelocationSymbol || "";
        }
    }
  getItems(): { [key: string]: string } {
    let bits = "";
    if (this.HasFunctions) bits += "[FN]";
    if (this.HasFilenames) bits += "[FL]";
    if (this.HasLineNumbers) bits += "[LN]";
    if (this.HasInlineFrames) bits += "[IN]";
    return {
      '1': `${this.ID}`,
      '2': `0x${this.Start.toString(16)}`,
      '3': `0x${this.Offset.toString(16)}`,
      '4': `0x${this.Limit.toString(16)}`,
      '5': `${this.File}`,
      '6': `${this.BuildID}`,
      '7': bits
    };
  }

  toString(): string {
    let bits = "";
    if (this.HasFunctions) bits += "[FN]";
    if (this.HasFilenames) bits += "[FL]";
    if (this.HasLineNumbers) bits += "[LN]";
    if (this.HasInlineFrames) bits += "[IN]";

    return `${this.ID}: 0x${this.Start.toString(16)}/0x${this.Limit.toString(
      16
    )}/0x${this.Offset.toString(16)} ${this.File} ${this.BuildID} ${bits}`;
  }
}

export class Location {
  ID: number = 0;
  Mapping: Mapping | null = null;
  Address: number = 0;
  Line: Line[] = [];
  IsFolded: boolean = false;

  private mappingIDX: number = 0;


  constructor(data?: Partial<Location>) {
    if (data) {
        this.ID = data.ID || 0;
        this.Mapping = data.Mapping ? new Mapping(data.Mapping) : null;
        this.Address = data.Address || 0;
        this.Line = (data.Line || []).map(ln => new Line(ln));
        this.IsFolded = data.IsFolded || false;
    }
}
  getItems(): { [key: string]: string } {
    let M =''
    if (this.Mapping) {
      M += `M=${this.Mapping.ID} `;
    }
    if (this.IsFolded) {
      M += "[F]";
    }
    const ss: { [key: string]: string } = {
      '1': this.ID.toString(),
      '2': `0x${this.Address.toString()}`,
      '3': M,
    };

    this.Line.forEach((line) => {
      if (line.Function) {
        ss['4'] = `${line.Function.Name}`;
        ss['5'] = `${line.Function.Filename} [line=${line.Line}; column=${line.Column}; startline=${line.Function.StartLine}]`;
        if (line.Function.Name !== line.Function.SystemName) {
         ss['6'] = `(${line.Function.SystemName})`;
        }
      } 
    });

    return ss;
  }

  toString(): string {
    const ss: string[] = [];
    let locStr = `${this.ID.toString().padStart(6)}: 0x${this.Address.toString(
      16
    )} `;

    if (this.Mapping) {
      locStr += `M=${this.Mapping.ID} `;
    }

    if (this.IsFolded) {
      locStr += "[F] ";
    }

    if (this.Line.length === 0) {
      ss.push(locStr);
    }

    this.Line.forEach((line) => {
      let lnStr = "??";
      if (line.Function) {
        lnStr = `${line.Function.Name} ${line.Function.Filename}:${line.Line}:${line.Column} s=${line.Function.StartLine}`;
        if (line.Function.Name !== line.Function.SystemName) {
          lnStr += `(${line.Function.SystemName})`;
        }
      }
      ss.push(locStr + lnStr);
      locStr = "             ";
    });

    return ss.join("\n");
  }
}

export class Line {
  Function: Function | null = null;
  Line: number = 0;
  Column: number = 0;

  private functionIDX: number = 0;

  constructor(data?: Partial<Line>) {
    if (data) {
        this.Function = data.Function ? new Function(data.Function) : null;
        this.Line = data.Line || 0;
        this.Column = data.Column || 0;
    }
}
}

export class Function {
  ID: number = 0;
  Name: string = "";
  SystemName: string = "";
  Filename: string = "";
  StartLine: number = 0;

  constructor(data?: Partial<Function>) {
    if (data) {
        this.ID = data.ID || 0;
        this.Name = data.Name || "";
        this.SystemName = data.SystemName || "";
        this.Filename = data.Filename || "";
        this.StartLine = data.StartLine || 0;
    }
}
}
