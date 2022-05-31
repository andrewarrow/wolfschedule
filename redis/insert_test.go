package redis

import "testing"

func TestInsertItem(t *testing.T) {

	InsertItem(1653948463, "Google Pixel 7 prototype reputedly shows up on eBay", "href1")
	InsertItem(1653948463, "Google Pixel 8 prototype reputedly shows up on eBay", "href2")
	InsertItem(1653950281, "Hurricane Agatha makes landfall in Mexico", "href2")
	InsertItem(1653953881, "Google Pixel 7 prototype reputedly shows up on eBay", "href2")
	InsertItem(1653953881, "Here's what's open and closed on Memorial Day 2022", "href2")
	InsertItem(1653957481, "Strawberries likely caused hepatitis A outbreak, FDA says", "href2")
	InsertItem(1653961081, "Widespread severe weather likely to end Memorial Weekend", "href2")
	InsertItem(1653964681, "Cloud forecast shows where skies will be perfect for meteor shower or possible meteor storm", "href2")
	InsertItem(1653968282, "Mona Lisa smeared with cream by disguised man", "href2")
	InsertItem(1653971882, "China falls short on big Pacific deal but finds smaller wins", "href2")
	InsertItem(1653975482, "Rafael Nadal and Novak Djokovic Meet at French Open Quarterfinals", "href2")

}
