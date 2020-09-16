// 缩略图地址
export const imageProxy = 'https://imageproxy.lizo.top'

export function replaceForImageProxy(oriUrl, paramStr) {
    // http(s)://host/key/key.jpg => https://imageproxy.host/paramStr/key/key.jpg
    var split = oriUrl.split("/", 2);
    return imageProxy + "/" + paramStr + "/" + split[1];
}