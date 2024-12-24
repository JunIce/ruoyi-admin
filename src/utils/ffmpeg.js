import { FFmpeg } from "@ffmpeg/ffmpeg";
import { fetchFile, toBlobURL } from "@ffmpeg/util";

function getFilenameFromUrl(url) {
  // 创建一个 URL 对象，这样可以更容易地处理各种 URL 组件
  var urlObj = new URL(url);
  // 获取 pathname（URL 中的路径部分），然后使用 split 方法将路径按斜杠分割成数组
  var pathArray = urlObj.pathname.split("/");
  // 获取数组的最后一项，这通常是文件名（如果存在的话）
  var filename = pathArray[pathArray.length - 1] || "";
  // 如果有查询参数或片段标识符，上面的代码会自动忽略它们
  return filename;
}

export async function ffmpegConvert(url, callback) {
  const ffmpeg = new FFmpeg();
  console.log(url);

  ffmpeg.on("log", ({ message: msg }) => {
    callback && callback(msg);
    console.log(msg);
  });
  await ffmpeg.load({
    coreURL: await toBlobURL("/ffmpeg-core.js", "text/javascript"),
    wasmURL: await toBlobURL("/ffmpeg-core.wasm", "application/wasm"),
    workerURL: await toBlobURL("/ffmpeg-core.worker.js", "text/javascript"),
  });

  let filename = getFilenameFromUrl(url);
  console.log("filename", filename);

  await ffmpeg.writeFile(filename, await fetchFile(url));
  console.log("writeFile success");
  await ffmpeg.exec(["-i", filename, "-vf", "scale=6000:400", "test.mp4"]);
  console.log("exec success");
  const data = await ffmpeg.readFile("test.mp4");
  console.log("readFile success", data);
  return data;
}
