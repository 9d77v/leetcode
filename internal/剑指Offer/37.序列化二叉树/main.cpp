#include <stdio.h>
#include <string>
using namespace std;
#include <queue>

//Definition for a binary tree node.
struct TreeNode
{
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

class Codec
{
public:
    // Encodes a tree to a single string.
    string serialize(TreeNode *root)
    {
        if (root == NULL)
        {
            return "";
        }
        string res = "";
        queue<TreeNode *> my_queue;
        my_queue.push(root);
        TreeNode *cur = new TreeNode(0);
        while (!my_queue.empty())
        {
            int len = my_queue.size();
            while (len--)
            {
                cur = my_queue.front();
                my_queue.pop();
                if (cur == NULL)
                {
                    res.push_back('$');
                }
                else
                {
                    res.append(to_string(cur->val));
                }
                res.push_back(',');
                if (cur != NULL)
                {
                    my_queue.push(cur->left);
                    my_queue.push(cur->right);
                }
            }
        }
        res.pop_back();
        return res;
    }

    // Decodes your encoded data to tree.
    TreeNode *deserialize(string data)
    {
        if (data.size() == 0)
        {
            return NULL;
        }
        int len = data.size();
        int i = 0;
        vector<TreeNode *> vec;
        while (i < len)
        {
            string str = "";
            while (i < len && data[i] != ',')
            {
                str.push_back(data[i]);
                i++;
            }
            if (str == "$")
            {
                vec.push_back(NULL);
            }
            else
            {
                int temp = std::stoi(str);
                TreeNode *cur = new TreeNode(temp);
                vec.push_back(cur);
            }
            i++;
        }
        int j = 1;
        for (int i = 0; j < vec.size(); i++)
        {
            if (vec[i] == NULL)
                continue;
            if (j < vec.size())
                vec[i]->left = vec[j++];
            if (j < vec.size())
                vec[i]->right = vec[j++];
        }
        return vec[0];
    }
};

// Your Codec object will be instantiated and called as such:
// Codec codec;
// codec.deserialize(codec.serialize(root));
