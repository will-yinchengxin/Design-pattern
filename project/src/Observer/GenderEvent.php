<?php
/**
 * Create By: Will Yin
 * Date: 2021/2/28
 * Time: 14:51
 **/
namespace Design\Observer;

class GenderEvent extends Event {

    /**
     * 这里触发观察者
     */
    public function trigger() {
        echo "接受到信号,要进行任务通知并更新!";
        // 执行更新
        $this->notify();
    }

}