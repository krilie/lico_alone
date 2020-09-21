import {getQuery} from "./ApiManage";

export const getVisitorPoints = () => {
    return getQuery("/manage/statistic/get_visitor_points"); // 高德地图key
}