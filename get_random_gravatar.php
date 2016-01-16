#!/usr/bin/env php
<?php
$random_md5 = md5(uniqid());
$size_in_pixels = 96;
$default_image = 'identicon';
$rating = "g";
file_put_contents("{$random_md5}.png", file_get_contents("http://www.gravatar.com/avatar/{$random_md5}?s={$size_in_pixels}&d={$default_image}&r={$rating}"));