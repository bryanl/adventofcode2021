const fs = require('fs');
const data = fs.readFileSync(('input' + __filename.slice(__dirname.length + 1, -4) + '.txt'), 'utf8').split('\r\n\r\n');

const enhancer = data[0];
const baseImage = data[1].split('\r\n');

function enhanceBorder(border) {
    switch (border) {
        case '.':
            return (enhancer[0] === '.') ? '.' : '#';
        case '#':
            return (enhancer[enhancer.length - 1] === '#') ? '#' : '.';
    }
}

function enhanceImage(image) {
    const newImage = [];
    for (let y = -1; y < image.length + 1; y++) {
        const newRow = [];
        for (let x = -1; x < image.length + 1; x++) {
            let address = '';
            for (let dy = -1; dy < 2; dy++) {
                for (let dx = -1; dx < 2; dx++) {
                    address += ((image[y + dy]?.[x + dx] ?? border) === '.') ? 0 : 1;
                }
            }
            address = parseInt(address, 2);
            newRow.push(enhancer[address]);
        }
        newImage.push(newRow.join(''));
    }
    border = enhanceBorder(border); // yes i'm using a global intentionally. sue me
    return newImage;
}

let border = '.';
let currentImage = baseImage;
for (let i = 0; i < 50; i++) { // EZ Clap no changes required
    currentImage = enhanceImage(currentImage, border);
}

let litCount = 0;
for (let y = 0; y < currentImage.length; y++) {
    for (let x = 0; x < currentImage[y].length; x++) {
        if (currentImage[y][x] === '#') litCount++;
    }
}

console.log(litCount);
