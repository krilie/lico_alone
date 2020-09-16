// 缩略图地址
export const imageProxy = 'https://imageproxy.lizo.top'

export function replaceForImageProxy(oriUrl, paramStr) {
    // http(s)://host/key/key.jpg => https://imageproxy.host/paramStr/key/key.jpg
    if (oriUrl.indexOf("minio.lizo.top") === -1)
        return imageProxy + "/" + paramStr + "/" + oriUrl;
    var split1 = oriUrl.split("//", 2);
    var split2 = split1[1].split("/", 2)
    return imageProxy + "/" + paramStr + "/" + split2[1];
}