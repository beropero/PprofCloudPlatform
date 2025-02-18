package template

var (
	EmailCodeTemplate = `
	<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>认证代码邮件</title>
    <style>
        /* 基线与移动端优化 */
        body {
            margin: 0;
            padding: 2vh 4vw;
            background: #f8f9ff;
            font-family: system-ui, -apple-system, sans-serif;
            line-height: 1.5;
        }

        /* 动态容器 */
        .main-container {
            width: min(95%, 680px);
            margin: 2vh auto;
            background: #fff;
            border-radius: clamp(8px, 1.2vw, 12px);
            box-shadow: 0 4px 12px rgba(28,75,128,0.06);
            padding: clamp(20px, 3vw, 32px);
        }

        /* 流体文字系统 */
        .brand {
            font-size: clamp(24px, 4vw, 30px);
            color: #2563eb;
            border-bottom: 1px solid #e0e7ff;
            padding-bottom: clamp(12px, 1.8vw, 16px);
        }

        /* 验证码动态容器 */
        .code-wrapper {
            --code-size: clamp(2.5rem, 12vw, 3.5rem);
            margin: 4vh 0;
            padding: clamp(12px, 2vw, 16px);
            background: #f6f7ff;
            border-radius: clamp(4px, 1vw, 8px);
        }

        /* 流动安全提示 */
        .security-alert {
            background: #fff5f6;
            border-color: #ff4b57;
            padding: clamp(12px, 2vw, 16px);
            margin: clamp(15px, 3vh, 25px) 0;
            display: grid;
            grid-template-columns: auto 1fr;
            gap: clamp(8px, 1.5vw, 12px);
        }

        @media (max-width: 480px) {
            .security-alert {
                grid-template-columns: 1fr;
                text-align: center;
            }
        }
    </style>
</head>
<body>
    <div class="main-container">
        <!-- 头部 -->
        <header style="margin-bottom: clamp(20px, 3vh, 28px)">
            <div class="brand">Pprof Cloud Platform</div>
        </header>

        <!-- 主要内容 -->
        <main>
            <p style="color: #4b5563; margin: 0 0 clamp(15px, 2.5vh, 20px)">
                您正在注册性能测试平台，请确认操作有效性：
            </p>

            <!-- 动态验证码区块 -->
            <div class="code-wrapper">
                <div style="color: #2563eb;
                          font-size: var(--code-size);
                          letter-spacing: 0.05em;
                          font-weight: 700;
                          text-align: center">
                    {{.code}}
                </div>
            </div>

            <!-- 自适安全提示 -->
            <div class="security-alert"
                 style="border-left: clamp(3px, 0.6vw, 4px) solid #ff4b57;
                        border-radius: clamp(4px, 0.8vw, 6px)">
                <div>
                    <p style="color: #6b7280; margin: 0;
                            font-size: clamp(13px, 2.5vw, 14px)">
                        验证码将在5分钟后失效<br class="desktop-br">
                        请勿将验证码告知第三方人员
                    </p>
                </div>
            </div>

            <!-- 响应式底部 -->
            <footer style="margin-top: clamp(20px, 4vh, 28px);
                         color: #9ca3af;
                         font-size: clamp(12px, 2.5vw, 13px)">
                <p style="margin: 0">系统自动发送，无需回复</p>
                <p style="margin: clamp(5px, 1vh, 8px) 0 0">ⓅPprof安全中心认证系统</p>
            </footer>
        </main>
    </div>
</body>
</html>

`
)
