package taskscheduler;

//https://leetcode.com/problems/task-scheduler/description/
// It looks like skill cool down
// 思路：
// 1 模拟选择不同的任务，
// 2.第一个选择数量最多的任务，
// 3.选择任务，在间隔足够的任务里，选最多的。
// 4.如果间隔都不足，就加idle

// 1. 统计每种任务的数量
// 2. 维护每种任务的数量，并找到当前数量最大的任务
// 3. 记录cpu使用的次数

// TODO should find a better solution

public class TaskScheduler {

    private int leastInterval(char[] tasks, int n) {

        int tasksNum = tasks.length;
        int[] taskCount = new int[26];
        int[] taskInterval = new int[26];
        int cpuNum = 0;
        for (char t: tasks) {
            taskCount[t - 'A'] += 1;
            taskInterval[t - 'A'] = 0;
        }

        while (tasksNum > 0) {
            int currentIndex = maxCountCharIndex(taskCount, taskInterval);
            if (currentIndex >= 0) {
                taskCount[currentIndex]--;
                for (int i = 0 ; i < taskInterval.length; i++) {
                    taskInterval[i]--;
                }
                taskInterval[currentIndex] = n;
                tasksNum--;
                cpuNum++;
            } else {
                for (int i = 0 ; i < taskInterval.length; i++) {
                    taskInterval[i]--;
                }
                cpuNum++;
            }
        }

        return cpuNum;
    }

    /**
     * find max count char , Interval is enough
     *
     * @param taskCount 代表任务数量
     * @param taskInterval  任务与之前的间隔
     * @return 返回数量最多，间隔足够的索引
     */
    private int maxCountCharIndex(int[] taskCount, int[] taskInterval) {
        int max = -1;
        int maxIndex = -1;
        for (int i = 0; i < taskCount.length; i++) {
            if (taskInterval[i] > 0) {
                continue;
            }
            if (taskCount[i] == 0) {
                continue;
            }
            if (taskCount[i] > max) {
                maxIndex = i;
                max = taskCount[i];
            }
        }
        return maxIndex;
    }

    public static void main(String[] args) {
        char[] task = new char[]{'A','A','A','B','B','B'};
        TaskScheduler so = new TaskScheduler();
        System.out.println(so.leastInterval(task, 2));
        System.out.println(so.leastInterval(task, 3));
    }
}
