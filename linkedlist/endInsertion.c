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
Insertion at end : similar to insertion at front
2 arguments > head_reference(pointer to pointer) & new_data
&& previous node MUST NOT be EMPTY*/
void insertEnd(struct Node** head_reference,int new_data){
	struct Node* new_node=(struct Node*)malloc(sizeof(struct Node));
	struct Node *last=*head_reference; // needed for traversal to last node
	new_node->data=new_data;
	new_node->next=NULL;
	//in case of empty list place head properly
	if(*head_reference==NULL){
		*head_reference=new_node;
		return;
	}
	while(last->next!=NULL){ //traverse to last list
		last=last->next;
	}
	last->next=new_node; //last node is the node we playing with
 	return;
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
insertEnd(&head,9); //push 9 at end
printList(head);
return 0;
}