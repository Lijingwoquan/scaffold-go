import { router } from "~/router"
import {
    showFullLoading,
    hideFullLoading
} from "~/composables/util"
import store from "./store"



router.beforeEach(async (to, from, next) => {
    //显示loading
    showFullLoading()
    await store.dispatch("getIndexInfo")
    next();
});

router.afterEach((to, from) => {
    //隐藏loading
    hideFullLoading()

});