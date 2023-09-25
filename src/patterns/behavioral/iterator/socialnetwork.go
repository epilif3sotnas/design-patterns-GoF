package iterator


type SocialNetwork interface {
	CreateFriendsIterator(string) ProfileIterator
	CreateCoworkersIterator(string) ProfileIterator
}