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
Insertion as frontal node : here we give reference of reference(pointers to pointers)
2 arguments > head_reference & new_data
*/
void push(struct Node** head_reference,int new_data){
	struct Node* new_node=(struct Node*)malloc(sizeof(struct Node));
	new_node->data=new_data;
	new_node->next=(*head_reference);
	(*head_reference)=new_node;
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
push(&head,5);
printList(head);
return 0;
}