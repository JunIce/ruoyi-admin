const path = require("path");
const fs = require("fs");
const { exec } = require("child_process");
const TARGET_FILE = path.join(__dirname, "../dist/index.html");

try {
  // 打开文件并将内容追加到文件末尾
  fs.appendFileSync(
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
    fs.writeFile(
      path.resolve(process.cwd(), "dist/info.txt"),
      str,
      function () {}
    );
  });
});
