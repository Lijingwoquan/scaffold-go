import axios from "~/axios.js";

export function getIndexInfo() {
    return new Promise((resolve, reject) => {
        axios.get("/base/index")
            .then(res => {
                resolve(res)
            })
            .catch(err => {
                reject(err)
            })
    })
}
