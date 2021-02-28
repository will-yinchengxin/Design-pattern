<?php
/**
 * Create By: Will Yin
 * Date: 2021/2/28
 * Time: 14:58
 **/
namespace Design\Observer;

class SchoolObserver implements Observe
{

    function update($event_info = '')
    {
        echo "学校接受到更新通知正在调整";
    }
}