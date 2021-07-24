const fs = require("fs");
const path = require("path");

const sharp = require("sharp");

function resizeImages(inFolder, outFolder, resizeWidth) {
  const start = Date.now();

  console.log("Reading files...");

  const files = fs.readdirSync(inFolder);

  if (fs.existsSync(outFolder)) {
    fs.rmSync(outFolder, {
      recursive: true,
      force: true,
    });
  }

  fs.mkdirSync(outFolder);

  console.log("Processing...");

  const tasks = [];

  let completed = 0;

  files.forEach((file) => {
    tasks.push(
      sharp(path.join(inFolder, file))
        .resize(resizeWidth, null, {
          fit: "contain",
        })
        .toFile(path.join(outFolder, file))
        .then(() => {
          if (completed++ % 100 === 0) {
            console.log(`${((completed / files.length) * 100).toFixed(0)}%`);
          }
        })
    );
  });

  return Promise.all(tasks);
}

module.exports = resizeImages;
