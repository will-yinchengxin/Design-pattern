<?php
/**
 * Create By: Will Yin
 * Date: 2021/2/28
 * Time: 14:06
 **/
namespace Design\Strategy;

class ManStrategy implements Strategy{

    function showAd()
    {
        echo "this is Ad for man";
    }

    function showCategory()
    {
        echo "this is category for man";
    }
}