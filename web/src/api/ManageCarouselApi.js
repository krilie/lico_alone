import {getQuery, postForm, postJson} from "./ApiManage";

export const manageGetCarouselList = () => {
    return getQuery("/manage/carousel/query");
}
export const manageDeleteCarouselById = (id) => {
    return postForm("/manage/carousel/delete_by_id", {carousel_id: id})
}

export const manageCreateCarousel = (data) => {
    // {
    //   "is_on_show": true,
    //   "message": "string",
    //   "url": "string"
    // }
    return postJson("/manage/carousel/create", {...data})
}

export const manageUpdateCarousel = (data) => {
    // {
    //   "id": "string",
    //   "is_on_show": true,
    //   "message": "string",
    //   "url": "string"
    // }
    return postJson("/manage/carousel/update", {...data})
}