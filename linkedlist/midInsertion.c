#include<stdio.h>
#include<stdlib.h>
// node declaration

struct Node{
	int data;
	struct Node *next;
};

// function to print nodes
void printList(struct Node *n){
	while(n!=NULL){
		printf("%d\n",n->data);
		n=n->next;
	}
}

/*
Insertion in Linked List
Insertion can be 3 ways :
>front
>somewhere in mid
>at end
*/

/*
Insertion somewhere in mid : we need the data to be pushed and reference from previous node
2 arguments > prev_node & new_data
&& previous node MUST NOT be EMPTY*/
void insert(struct Node* prev_node,int new_data){
	if(prev_node==NULL){
		printf("\n previous node argument is empty, BAD LOCATION");
		return;
	}
	struct Node* new_node=(struct Node*)malloc(sizeof(struct Node));
	new_node->data=new_data;
	new_node->next=prev_node->next;
	prev_node->next=new_node;
	}
// lets make 3 nodes
int main(){
	struct Node* head=NULL;
	struct Node* second=NULL;
	struct Node* third=NULL;

	// allocate 3 blank nodes

	head=(struct Node*)malloc(sizeof(struct Node));
	second=(struct Node*)malloc(sizeof(struct Node));
	third=(struct Node*)malloc(sizeof(struct Node));
// Assign head
	head->data=1234;
	head->next=second;
// Assign second
	second->data=5678;
	second->next=third;
// Assign third
	third->data=91011;
	third->next=NULL;
// Print the damn list.
insert(head->next,6); //push 6 next to head
printList(head);
return 0;
}