#include<iostream>
#include<vector>
#include<algorithm>
using namespace std;
int main(){
    int num;
    vector<int> nums;
    while(cin>>num){
        nums.push_back(num);
        if(getchar()=='\n'){
            reverse(nums.begin(),nums.end());
            for(int i=0;i<nums.size()-1;i++)
                cout<<nums[i]<<" ";
            cout<<nums.back()<<"\n";
            nums.clear();
        }
    }
    return 0;
}