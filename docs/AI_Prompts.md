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
│   ├── handlers/   # 后端处理函数
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

## 第二阶段：聊天界面实现（阶段 2）

### 任务 1：实现后端处理函数

在 `src/handlers/` 目录中创建以下处理函数：

1. **应用处理函数** (`app.go`)：
   - `getAppInfo`: 返回应用信息
   - `verifyAppUser`: 用户认证
   - `saveAppUserInfo`: 更新用户资料
   - `updateAppUserPassword`: 更改用户密码
   - `getAppUserActionLogList`: 获取用户操作日志

2. **平台处理函数** (`platform.go`)：
   - `getGitPlatformList`: 返回 Git 平台列表
   - `saveGitPlatform`: 添加/更新 Git 平台配置

3. **仓库处理函数** (`repo.go`)：
   - `getGitRepoList`: 返回平台的仓库列表
   - `getGitRepoInfo`: 返回仓库详细信息
   - `saveGitRepo`: 添加/更新仓库配置

4. **问题处理函数** (`issue.go`)：
   - `getGitIssueList`: 返回仓库的问题列表
   - `getGitIssueCommentList`: 返回问题的评论列表
   - `getGitIssueParticipantList`: 返回问题的参与者列表
   - `saveGitIssueComment`: 添加新评论到问题

### 任务 2：实现前端组件

在 `src/ui/src/components/` 目录中创建以下组件，并按功能组织：

1. **应用组件** (`app/`)：
   - `AppAuthPage.svelte`: 认证页面，包含登录表单
   - `AppMainPage.svelte`: 主应用布局
   - `AppMainPageHeader.svelte`: 应用头部，包含用户菜单
   - `AppMainPageFooter.svelte`: 应用底部
   - `AppLockDialog.svelte`: 应用锁定模态对话框
   - `AppUserActionLogPanel.svelte`: 用户操作日志面板

2. **仓库组件** (`repo/`)：
   - `GitPlatformList.svelte`: Git 平台列表
   - `GitPlatformAddDialog.svelte`: 添加 Git 平台对话框
   - `GitRepoList.svelte`: 仓库列表
   - `GitRepoAddDialog.svelte`: 两步添加仓库对话框
   - `GitRepoPanel.svelte`: 仓库面板，包含平台和仓库列表

3. **问题组件** (`issue/`)：
   - `GitIssueList.svelte`: 问题列表
   - `GitIssueAddDialog.svelte`: 添加问题对话框
   - `GitIssuePanel.svelte`: 问题面板
   - `GitIssueCommentList.svelte`: 评论列表
   - `GitIssueCommentPanel.svelte`: 评论面板
   - `GitIssueParticipantList.svelte`: 参与者列表

4. **用户组件** (`user/`)：
   - `UserProfileDialog.svelte`: 用户资料对话框
   - `UserPasswordDialog.svelte`: 用户密码对话框

5. **工具组件** (`lib/components/`)：
   - `Dialog.svelte`: 可复用的基础对话框组件
   - `Notification.svelte`: 通知显示组件

### 任务 3：实现状态管理

在 `src/ui/src/lib/store.js` 中实现：

1. 应用状态（用户信息、选中项）
2. UI 状态（对话框可见性、通知）
3. 数据列表（平台、仓库、问题、评论、参与者）
4. 操作日志
5. 状态更新辅助函数

### 任务 4：实现桥接层

在 `src/ui/src/lib/bridge.js` 中实现：

1. 通用函数调用机制用于 Go 处理函数
2. 参数序列化和返回值反序列化
3. 错误处理和通知显示
4. 所有后端处理函数的特定包装函数

### 任务 5：UI/UX 改进

1. 使用 Tailwind CSS 实现响应式布局：
   - 固定比例宽度（仓库面板 1/6，问题面板 5/6）
   - 组件间的视觉分隔线
   - 当前用户评论右侧显示

2. 增强用户体验：
   - 操作日志面板用于跟踪用户活动
   - 模态对话框替代独立页面
   - 键盘支持（Esc 关闭对话框）
   - 加载状态和错误处理

### 任务 6：代码现代化

1. 将所有 Svelte 组件转换为 TypeScript 并使用 Svelte 5 runes：
   - 使用 `$props()` 处理组件属性
   - 使用 `$state()` 处理响应式状态
   - 使用 `$effect()` 处理副作用

2. 更新回调命名约定为 `onFunction` 格式

3. 用 `$props` 回调替换已弃用的 `createEventDispatcher`

### 任务 7：组件重构

1. 按功能重新组织前端组件：
   - `app/`: 应用级组件
   - `chat/`: 聊天面板组件
   - `issue/`: 问题相关组件
   - `repo/`: 仓库相关组件
   - `user/`: 用户相关组件

2. 更新导入路径以反映新结构

### 任务 8：Bug 修复

1. 修复 `AppMainPageHeader.svelte` 中未定义的登出函数
2. 解决对话框背景样式问题
3. 修正组件属性处理
4. 修复后端函数命名不一致问题

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

### 组件架构

前端组件应按功能组织：
- `app/`: 应用级组件（认证、主布局、对话框）
- `chat/`: 聊天面板组件
- `issue/`: 问题相关组件（列表、面板、评论）
- `repo/`: 仓库相关组件（平台、仓库）
- `user/`: 用户相关组件（资料、密码管理）

### 状态管理

应用状态应通过 Svelte stores 管理，明确分离关注点：
- 应用状态（用户信息、选中项）
- UI 状态（对话框可见性、通知）
- 数据列表（平台、仓库、问题、评论、参与者）
- 操作日志

### 桥接层

前后端通信应通过 JavaScript 桥接层处理：
- 为 Go 处理函数提供通用函数调用机制
- 处理参数序列化和返回值反序列化
- 实现错误处理和通知显示
- 为所有后端处理函数暴露特定包装函数

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