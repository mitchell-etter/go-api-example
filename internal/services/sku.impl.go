package services

func NewSkuService() SkuService {
	return &skuServiceImpl{}
}

type skuServiceImpl struct {
}

func (svc *skuServiceImpl) GetAll() []SkuItem {
	return []SkuItem{}
}
