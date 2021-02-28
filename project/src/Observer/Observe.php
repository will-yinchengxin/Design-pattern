<?php
/**
 * Create By: Will Yin
 * Date: 2021/2/28
 * Time: 14:47
 **/
namespace Design\Observer;

interface Observe {

    /**
     * @param string $event_info
     */
    function update ($event_info = '');

}