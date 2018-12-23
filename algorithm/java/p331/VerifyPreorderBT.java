package p331;


// https://leetcode.com/problems/verify-preorder-serialization-of-a-binary-tree/description/
// 提示使用stack；
// java linkedList; push pop empty(peek)
//  1. 初始化一个stack，放一个元素
//  2. 访问一个字符， 判断stack是否为空，如果为空就异常，不为空，就是有位置放这个节点，pop；判断字符为普通字符，则压栈两个位置，就是这个节点可以继续放下两个位置；如果字符为'#',不压栈
//  3. 字符访问完毕，即为结束。
// 由于这里只是检查是否符合，不用重建，只要记录位置数量即可。
//
class Solution {
    public boolean isValidSerialization(String preorder) {
        String[] chars = preorder.split(",");
        int pos_num = 1;
        if (chars.length ==0) {
            return false;
        }

        for (int i = 0; i < chars.length; i += 1) {
            if (pos_num==0) {
                return false;
            }

            pos_num--;

            if (!chars[i].equals("#")) {
                pos_num += 2;
            }
        }
        if (pos_num != 0) {
            return false;
        }
        return true;
    }
}
public class VerifyPreorderBT {

    public static void main(String[] args) {
        System.out.println(new Solution().isValidSerialization("9,#,3"));
    }
}
