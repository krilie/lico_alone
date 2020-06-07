import {getQuery, postForm, postMultiForm} from "./ApiManage";

// 管理员获取文件列表
// page_num page_size key_name_like bucket_name_like url_like user_id biz_type content_type created_at_begin created_at_end
export const manageGetFilePage = (querys) => {
    return getQuery("/manage/file/query", {...querys});
}

export const manageDeleteFile = (fileId) => {
    return postForm("/manage/file/delete", {file_id: fileId})
}

export const manageUpdateFile = (file) => {
    return postMultiForm("/manage/file/upload", {file: file})
}