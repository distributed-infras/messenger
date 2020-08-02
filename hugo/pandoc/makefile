convert:
	pandoc --extract-media ./Images hyperchat.odt -o hyperchat.md
images:
	sed 's+./Images+/Images/design+g' hyperchat.md > hyperchat-2.md
anchors:
	cat design.md | perl -ne 's/\[\]\{(#anchor-\d+)\}/# []($$1) /g;print $$_;' > design2.md
image2:
	cat design2.md | perl -ne 's/\/Images\//\/messenger\/Images\//g;print $$_;' > design3.md

