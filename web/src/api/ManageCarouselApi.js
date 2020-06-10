import {getQuery} from "./ApiManage";

export const manageGetCarouselList = () => {
    return getQuery("/manage/carousel/query");
}