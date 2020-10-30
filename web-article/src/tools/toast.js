import {toast} from "react-toastify";

export const ToastWarn = (msg) => toast.warn(msg);
export const ToastErr = (msg) => toast.error(msg);
export const ToastNormal = (msg) => toast.info(msg);