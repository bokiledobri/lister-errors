.PHONY: publish
publish:
	git tag ${TAG} \
		&& git push origin ${TAG} \
		&& GOPROXY=proxy.golang.org go list -m github.com/bokiledobri/lister-errors@${TAG}
