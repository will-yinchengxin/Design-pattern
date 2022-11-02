<?php
/**
 * Create By: Will Yin
 * Date: 2021/2/28
 * Time: 14:59
 **/
namespace Design\Observer;


class ShopObserver implements Observe {

    function update($event_info = '')
    {
        echo "商贸接受到更新通知正在调整";
    }
}