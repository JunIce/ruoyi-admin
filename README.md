# ruoyi-admin

基于 ruoyi vue3、element-plus 的界面优化后台管理系统

## Usage

- Vue
- Element-Plus
- Vite
- ffmpeg-core.wasm

## Install Dependencies

```bash
yarn install or npm install
```

## Run

```bash
yarn dev or npm run dev
```

## Build

```bash
yarn build:prod or npm run build:prod
```

## 功能

- ffmpeg 前端视频转码播放

  - 使用 ffmpeg-core.js 和 ffmpeg-core.wasm 进行视频转码
  - 使用 ffmpeg-core.worker.js 进行视频转码
  - 目前在前端转换视频<b>比较慢</b>，只可以做一种解决方案，如果需要做视频转码，最好是后端转码

- 视频转码进度条

## References

[RuoYi-Vue3](https://github.com/yangzongzhuan/RuoYi-Vue3.git)
