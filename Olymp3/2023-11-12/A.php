<?php
$d = intval(readline());

$solutions = [];

$temp_x = 0;
$temp_y = floor($d * sqrt(2));

while ($temp_x <= $d) {
    if ($temp_x ** 2 + $temp_y ** 2 == 2 * $d ** 2) {
        $solutions[] = [$temp_x, $temp_y];
        $temp_y -= 1;
    } elseif ($temp_x ** 2 + $temp_y ** 2 < 2 * $d ** 2) {
        $temp_x += 1;
    } else {
        $temp_y -= 1;
    }
}

echo count($solutions) . PHP_EOL;
foreach ($solutions as $s) {
    echo implode(' ', $s) . PHP_EOL;
}
?>
