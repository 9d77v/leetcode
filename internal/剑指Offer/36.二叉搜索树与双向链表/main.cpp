#include <stdio.h>

class Node
{
public:
    int val;
    Node *left;
    Node *right;

    Node() {}

    Node(int _val)
    {
        val = _val;
        left = NULL;
        right = NULL;
    }
    Node(int _val, Node *_left, Node *_right)
    {
        val = _val;
        left = _left;
        right = _right;
    }
};

class Solution
{
public:
    Node *tail, *head;
    Node *treeToDoublyList(Node *root)
    {
        if (root == NULL)
        {
            return NULL;
        }
        dfs(root);
        head->left = tail;
        tail->right = head;
        return head;
    }

    void dfs(Node *cur)
    {
        if (cur == NULL)
        {
            return;
        }
        dfs(cur->left);
        if (tail != NULL)
        {
            tail->right = cur;
        }
        else
        {
            head = cur;
        }
        cur->left = tail;
        tail = cur;
        dfs(cur->right);
    }
};
