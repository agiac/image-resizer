const fs = require("fs");

const commandLineArgs = require("command-line-args");

const resizeImages = require("./resizeImages");

const fatal = (message) => {
  console.error(`ERROR: ${message}`);
  process.exit(1);
};

const options = commandLineArgs([
  {
    name: "in",
    type: String,
    defaultValue: "",
  },
  { name: "out", type: String, defaultValue: "" },
  { name: "width", type: Number, defaultValue: 0 },
]);

const { in: inFolder, out: outFolder, width: resizeWidth } = options;

if (inFolder === "") {
  fatal("Input folder is undefined");
}

if (outFolder === "") {
  fatal("Output folder is undefined");
}

if (resizeWidth === 0) {
  fatal("Width is undefined");
}

if (resizeWidth < 0) {
  fatal("Width must be a positive number");
}

if (!fs.existsSync(inFolder)) {
  fatal(`Input folder '${inFolder}' does not exist`);
}

const start = Date.now();

resizeImages(inFolder, outFolder, resizeWidth)
  .then(() => {
    console.log(`Done in ${((Date.now() - start) / 1000).toFixed()} seconds.`);
  })
  .catch(console.error);
