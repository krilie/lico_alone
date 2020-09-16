// 缩略图地址
export const imageProxy = 'https://imageproxy.lizo.top'

export function replaceForImageProxy(oriUrl, paramStr) {
    // http(s)://host/key/key.jpg => https://imageproxy.host/paramStr/key/key.jpg
    var myURL = new URL(oriUrl);
    var split = oriUrl.split(myURL.host + "/", 2);
    return imageProxy + "/" + paramStr + "/" + split[1];
}