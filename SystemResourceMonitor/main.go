package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"syscall"
)

func getCPUUsage() (float64, error) {
	data, err := ioutil.ReadFile("/proc/stat")
	if err != nil {
		return 0, err
	}

	lines := strings.Split(string(data), "\n")
	if len(lines) < 1 {
		return 0, fmt.Errorf("unexpected format in /proc/stat")
	}

	fields := strings.Fields(lines[0])
	if len(fields) < 5 {
		return 0, fmt.Errorf("unexpected format in /proc/stat")
	}

	user, err := strconv.ParseUint(fields[1], 10, 64)
	if err != nil {
		return 0, err
	}
	nice, err := strconv.ParseUint(fields[2], 10, 64)
	if err != nil {
		return 0, err
	}
	system, err := strconv.ParseUint(fields[3], 10, 64)
	if err != nil {
		return 0, err
	}
	idle, err := strconv.ParseUint(fields[4], 10, 64)
	if err != nil {
		return 0, err
	}

	total := user + nice + system + idle
	usage := float64(user+nice+system) / float64(total) * 100

	return usage, nil
}

func getMemoryUsage() (float64, error) {
	data, err := ioutil.ReadFile("/proc/meminfo")
	if err != nil {
		return 0, err
	}

	lines := strings.Split(string(data), "\n")
	if len(lines) < 2 {
		return 0, fmt.Errorf("unexpected format in /proc/meminfo")
	}

	totalFields := strings.Fields(lines[0])
	if len(totalFields) < 2 {
		return 0, fmt.Errorf("unexpected format in /proc/meminfo")
	}
	total, err := strconv.ParseUint(totalFields[1], 10, 64)
	if err != nil {
		return 0, err
	}

	availableFields := strings.Fields(lines[1])
	if len(availableFields) < 2 {
		return 0, fmt.Errorf("unexpected format in /proc/meminfo")
	}
	available, err := strconv.ParseUint(availableFields[1], 10, 64)
	if err != nil {
		return 0, err
	}

	used := total - available
	usage := float64(used) / float64(total) * 100

	return usage, nil
}

func getDiskUsage(path string) (float64, error) {
	var stat syscall.Statfs_t
	if err := syscall.Statfs(path, &stat); err != nil {
		return 0, err
	}

	total := stat.Blocks * uint64(stat.Bsize)
	free := stat.Bfree * uint64(stat.Bsize)
	used := total - free
	usage := float64(used) / float64(total) * 100

	return usage, nil
}

func main() {
	cpuUsage, err := getCPUUsage()
	if err != nil {
		fmt.Printf("Error getting CPU usage: %v\n", err)
		return
	}

	memUsage, err := getMemoryUsage()
	if err != nil {
		fmt.Printf("Error getting memory usage: %v\n", err)
		return
	}

	diskUsage, err := getDiskUsage("/")
	if err != nil {
		fmt.Printf("Error getting disk usage: %v\n", err)
		return
	}

	fmt.Printf("CPU Usage: %.2f%%\n", cpuUsage)
	fmt.Printf("Memory Usage: %.2f%%\n", memUsage)
	fmt.Printf("Disk Usage: %.2f%%\n", diskUsage)
}
