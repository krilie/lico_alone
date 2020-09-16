// 缩略图地址
export const imageProxy = 'https://imageproxy.lizo.top'

export function replaceForImageProxy(oriUrl, paramStr) {
    // http(s)://host/key/key.jpg => https://imageproxy.host/paramStr/key/key.jpg
    var split1 = oriUrl.split("//",2);
    var split2 = split1[1]("/",2)
    return imageProxy + "/" + paramStr + "/" + split2[1];
}