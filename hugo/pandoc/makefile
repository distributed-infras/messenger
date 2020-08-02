convert:
	pandoc --extract-media ./Images hyperchat.odt -o hyperchat.md
images:
	sed 's+./Images+/Images/design+g' hyperchat.md > hyperchat-2.md
anchors:
	cat design.md | perl -ne 's/\[\]\{(#anchor-\d+)\}/# []($$1) /g;print $$_;' > design2.md
