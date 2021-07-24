const fs = require("fs");
const path = require("path");

const resizeImages = require("./resizeImages");

describe("Resize images", () => {
  it("shoudl correctly resize the provieded images", async () => {
    const inFolder = path.join(__dirname, "..", "..", "test-images");
    const outFolder = "/tmp/image-resizer-output";
    const resizeWidth = 500;

    await resizeImages(inFolder, outFolder, resizeWidth);

    const inFiles = fs.readdirSync(inFolder);
    const outFiles = fs.readdirSync(outFolder);

    expect(inFiles.length).toBe(outFiles.length);
  });
});
