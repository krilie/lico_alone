import {getQuery, postForm, postJson} from "./ApiManage";

export const manageGetArticleById = (id) => {
    return getQuery("/manage/article/get_by_id", {id});
}

export const manageDeleteArticle = (id) => {
    return postForm("/manage/article/delete", {article_id: id})
}

export const manageQueryArticle = (searchKey, pageNum, pageSize) => {
    // search_key page_num page_size
    return getQuery("/manage/article/query", {search_key: searchKey, page_num: pageNum, page_size: pageSize})
}

export const manageUpdateArticle = (data) => {
    // {
    //     "content": "string",
    //     "description": "string",
    //     "id": "string",
    //     "picture": "string",
    //     "sort": 0,
    //     "title": "string"
    // }
    return postJson("/manage/article/update", {...data})
}

export const manageCreateArticle = (data) => {
    // {
    //     "content": "string",
    //     "description": "string",
    //     "picture": "string",
    //     "sort": 0,
    //     "title": "string"
    // }
    return postJson("/manage/article/create", {...data})
}