<?php
// show_logs.php
$logFile = '/var/www/pensionmanagement/storage/logs/laravel.log'; // Path to your Nginx error log

// Check if the log file exists and is readable
if (file_exists($logFile) && is_readable($logFile)) {
    // Execute the tail command to get the last 100 lines
    $output = shell_exec('tail -n 100 ' . escapeshellarg($logFile));
    echo '<pre>' . htmlspecialchars($output) . '</pre>';
} else {
    echo 'Log file not found or not readable';
}
?>
