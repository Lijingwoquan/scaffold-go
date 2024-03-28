import axios from "~/axios.js";
import { toast } from '~/composables/util'

export function deleteImg(name) {
    return new Promise((resolve, reject) => {
        axios.delete("/manager/deleteImg", {
            params: {
                name: name
            }
        })
            .then(res => {
                toast("删除图片成功", "success")
                resolve(res);
            })
            .catch(err => {
                toast("删除图片失败","error")
                reject(err);
            });
    })
}


