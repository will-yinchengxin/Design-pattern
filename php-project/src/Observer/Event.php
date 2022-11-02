<?php
/**
 * Create By: Will Yin
 * Date: 2021/2/28
 * Time: 14:44
 **/
namespace Design\Observer;

abstract class Event {

    /**
     * @var array $observer
     */
    private $observer = [];

    /**
     * 添加观察者
     */
    public function addObserver(Observe $observe) {
        $this->observer[] = $observe;
    }

    /**
     * @监听事件
     */
    public function notify() {
        foreach ($this->observer as $observer) {
            $observer->update();
        }
    }

}