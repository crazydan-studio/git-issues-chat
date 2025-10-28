# Git Issues Chat - AI 实现提示

本文档包含了实现 Git Issues Chat 桌面应用程序的完整、有组织的提示和要求。这些提示专为 AI 助手设计，以系统化的方式执行项目实现。

## 项目概述

设计并实现一个名为 `Git Issues Chat` 的桌面聊天应用程序，该应用使用 GitHub 仓库作为聊天平台，仓库作为聊天室，Issues 作为话题，Comments 作为消息。这使得通过聊天界面实现实时问题处理和协作成为可能。

## 技术栈

- **后端**: Golang 用于核心应用逻辑、数据管理和 GitHub API 集成
- **前端**: Svelte 配合 Tailwind CSS 用于用户界面
- **UI 框架**: Go-WebUI 用于桌面应用程序部署
- **数据库**: SQLite 用于本地数据存储
- **构建系统**: Vite 用于 UI 打包

## 项目结构

```
git-issues-chat/
├── build/          # 构建脚本
│   ├── bin/        # 运行时脚本
│   └── build.sh    # 主构建脚本
├── deps/           # 第三方依赖
├── dist/           # 构建输出
│   ├── bin/        # 可执行文件
│   ├── ui/         # UI 构建输出
│   ├── data/       # 应用运行时数据
│   └── chat.sh     # 运行时脚本
├── docs/           # 文档
│   └── stages/     # 阶段开发文档
├── src/            # 源代码
│   ├── main.go     # 应用入口点
│   └── ui/         # Svelte UI 项目
└── VERSION         # 应用版本
```

## 第一阶段：项目基础（阶段 0）

### 任务 1：初始化项目结构

按照上述要求创建项目目录结构：

1. 初始化 git 仓库
2. 创建 VERSION 文件，版本为 "0.1.0"
3. 添加 Apache 2.0 LICENSE 文件
4. 设置适当的 .gitignore 以排除 dist/ 和 deps/ 目录

### 任务 2：设置 Golang 项目

创建 `src/main.go`，包含以下功能：

1. 命令行参数解析：
   - `--data-path`: 数据目录路径（必需）
   - `--ui-path`: UI 目录路径（必需）
   - `--browser`: 使用的浏览器（firefox, chrome, edge）（必需）
   - `--version`: 打印版本信息（"Git Issues Chat v0.1.0"）

2. 数据库初始化函数 `initAppEnv`：
   - 接受 dataPath 参数
   - 如果不存在则创建数据目录
   - 在 `{dataPath}/chat.db` 初始化 SQLite 数据库
   - 返回数据库对象

3. WebUI 集成：
   - 创建 WebUI 窗口
   - 设置根文件夹为 ui-path
   - 使用指定浏览器显示 index.html
   - 等待窗口关闭

4. 依赖：
   - github.com/mattn/go-sqlite3 用于 SQLite 支持
   - github.com/webui-dev/go-webui/v2 用于 WebUI 集成

### 任务 3：设置 Svelte + Tailwind CSS UI

在 `src/ui` 中初始化 Svelte 项目：

1. Tailwind CSS 集成：
   - 安装 tailwindcss, postcss, autoprefixer
   - 配置 tailwind.config.js 和 postcss.config.js
   - 在 app.css 中添加 Tailwind 指令

2. 欢迎页面要求：
   - 全屏居中布局
   - 文本："👏 Hello, WebUI Desktop!" 字体大小 72px
   - Body 应占据整个视口

3. 构建配置：
   - 设置输出目录为 `../../dist/ui`
   - 在 index.html 中包含 webui.js 脚本

4. 包配置：
   - 设置包名为 "git-issues-chat-ui"
   - 设置版本与项目版本匹配

### 任务 4：实现构建系统

创建 `build/build.sh` 脚本：

1. 从 VERSION 文件读取版本
2. 更新版本信息：
   - src/main.go（用于 --version 标志）
   - src/ui/package.json
3. 清理 dist 目录
4. 使用 Tailwind CSS 支持构建 UI
5. 获取 go-webui 依赖：
   - 如果不存在则克隆到 deps/go-webui
   - 如果存在则更新
6. 构建 Go 应用程序：
   - 发布版本：dist/bin/chat 带优化标志（-s -w）
   - 调试版本：dist/bin/chat-debug 带调试标志
7. 复制 build/bin/chat.sh 到 dist/chat.sh
8. 包含所有关键步骤的错误检查

### 任务 5：实现运行时脚本

创建 `build/bin/chat.sh` 脚本：

1. 命令行参数：
   - `--debug`: 运行调试版本
   - `--browser`: 指定浏览器（firefox, chrome, edge）
   - `--help`: 显示使用信息

2. 路径管理：
   - 使用相对于脚本位置的路径
   - DATA_PATH: {script_dir}/data
   - UI_PATH: {script_dir}/ui
   - EXECUTABLE: {script_dir}/bin/chat 或 chat-debug

3. 目录管理：
   - 如果不存在则自动创建数据目录
   - 验证 UI 目录存在

4. 权限检查：
   - 数据目录：可读可写
   - UI 目录和文件：可读
   - 可执行文件：可执行

5. 环境：
   - 设置 WEBKIT_DISABLE_DMABUF_RENDERER=1
   - 向应用程序传递正确参数

### 任务 6：文档

创建/更新文档文件：

1. `docs/stages/stage-0.md`：
   - 记录所有已完成任务
   - 记录遇到的问题和解决方案
   - 包含 Tailwind CSS 字体大小问题的技术细节：
     * 问题：由于 Tailwind v4 清理，预定义类不工作
     * 解决方案：使用任意值表示法（text-[72px]）
   - 记录文件组织变更
   - 记录权限和目录创建变更

2. `docs/design.md`：
   - 更新项目结构图
   - 记录 Tailwind CSS 处理的技术说明
   - 更新运行时脚本信息

3. `README.md`：
   - 更新项目结构
   - 更新使用说明
   - 记录先决条件（Node.js 20+）

### 任务 7：版本控制

提交所有更改并附上适当的提交信息，包括：
- 项目结构初始化
- 核心应用程序实现
- UI 实现
- 构建系统实现
- 文档更新
- 错误处理改进

## 技术要求和约束

### Tailwind CSS 字体大小处理

由于 Tailwind CSS v4 的内容基础清理，预定义的字体大小类可能不会包含在构建中。为了可靠的字体大小控制，使用任意值表示法（例如 `text-[72px]`）而不是预定义类（例如 `text-8xl`）。

### 数据目录管理

应用程序不应自动创建数据目录。此责任应转移到运行时脚本，该脚本应在数据目录不存在时自动创建，从而提供更好的关注点分离。

### 错误处理

所有关键操作都应包含适当的错误处理：
- Go-webui 获取失败
- 应用程序可执行文件构建失败
- 浏览器窗口打开失败
- 文件权限和存在性检查

### 路径管理

所有脚本都应使用相对于其自身位置的路径，以获得更好的可移植性，而不是绝对路径或复杂的路径解析逻辑。

## 质量保证

1. 所有构建应无错误完成
2. 生成的可执行文件应正确运行
3. UI 应正确显示，字体大小正确
4. 所有文档应准确且最新
5. Git 提交应具有原子性且文档完整
6. 错误消息应清晰且有帮助
7. 帮助文档应全面

## 验证步骤

1. 成功运行构建脚本
2. 验证 dist 目录内容：
   - bin/chat 和 bin/chat-debug 可执行文件
   - ui/ 目录与构建资产
   - chat.sh 脚本
   - data/ 目录
3. 使用各种选项测试运行时脚本
4. 验证帮助功能是否正常工作
5. 检查缺失文件/目录的错误处理
6. 确认所有文档准确完整