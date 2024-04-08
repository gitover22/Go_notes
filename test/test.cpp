#include<iostream>
#include<vector>
using namespace std;
class Solution {
public:
    int removeDuplicates(vector<int>& nums) {
        int len = nums.size();
        if(len == 1) return 1;
        int value =nums[0]; // 记录每一连续相同值
        int count_value = 1; // 记录值为value的个数
        int ptr_i = 0,flag=0;
        for(int i=1;i<len;i++){
                if(nums[i] != value || i == len-1 ){
                        if(count_value >= 2){
                            if(i == len-1 && ptr_i ==0 && nums[i] == value) flag = 1;
                            nums[ptr_i] = nums[ptr_i+1] =value;
                            ptr_i += 2;
                            if(flag == 1) return ptr_i;
                            if(i == len-1 && nums[i] == value) return ptr_i;
                        }else{
                            // count_value == 1的情况
                            nums[ptr_i] =value;
                            ptr_i++; // 慢指针移动
                            if(nums[i] == value && i == len-1)
                            {
                                nums[ptr_i] = value;
                                return ptr_i+1;
                            }
                        }
                        if( i== len-1 && nums[i] != value){
                            nums[ptr_i] = nums[i];
                            return ptr_i+1;
                        }
                        value = nums[i]; // 赋新值
                        count_value = 1; //重置
                }else{
                        count_value++;
                }
        }
        return ptr_i+1;
    }
};

int main()
{
    vector<int  >vec={1,1,1,2,2,3} ;
    Solution sol;
    int num = sol.removeDuplicates(vec);
    std::cout<<num<<std::endl;
    for(int i=0;i<num;i++){
        std::cout<<vec[i]<<" "<<std::endl;
    }

    return 0;
}
