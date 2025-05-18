import { join, resolve, dirname } from "path";
import { appendFileSync, writeFile } from "fs";
import { exec } from "child_process";
import { fileURLToPath } from "url";
// 获取当前文件的路径
const __filename = fileURLToPath(import.meta.url);
// 获取当前文件所在目录的路径
const __dirname = dirname(__filename);

const TARGET_FILE = join(__dirname, "../dist/index.html");

try {
  // 打开文件并将内容追加到文件末尾
  appendFileSync(
    TARGET_FILE,
    `<script>console.log('bundletime: ${String(new Date())}');</script>`
  );
  console.log("\x1b[32m%s\x1b[0m", "append bundletime success");
} catch (error) {
  console.error(`append bundletime error: ${error}`);
}

exec("git log -n 10 ", function (err, stdout, stderr) {
  if (err) return;
  exec("git rev-parse --abbrev-ref HEAD", function (err1, stdout1, stderr1) {
    let str = `build branch: ${stdout1} \n\n\n git log: \n${stdout} \n`;
    writeFile(resolve(process.cwd(), "dist/info.txt"), str, function () {});
  });
});
