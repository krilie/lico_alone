export default class FilePathHelper {
    /*获取这个路径的上一层路径*/
    static GetUpperPath(path) {
        var split = path.split('\\');
        if (split.length <= 0)
            return "";
        else if (split.length === 1)
            return "";
        else {
            return split.slice(0, split.length - 1).join("\\");
        }
    }
}