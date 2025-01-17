package v1alpha1

import (
	"github.com/metallb/metallb-operator/api/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// ConvertTo converts this AddressPool to the Hub version (vbeta1).
func (src *AddressPool) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1beta1.AddressPool)
	dst.Spec.Protocol = src.Spec.Protocol
	dst.Spec.Addresses = make([]string, len(src.Spec.Addresses))
	copy(dst.Spec.Addresses, src.Spec.Addresses)
	*dst.Spec.AutoAssign = *src.Spec.AutoAssign
	dst.Spec.BGPAdvertisements = make([]v1beta1.BgpAdvertisement, 0)
	for i, adv := range src.Spec.BGPAdvertisements {
		*dst.Spec.BGPAdvertisements[i].AggregationLength = *adv.AggregationLength
		*dst.Spec.BGPAdvertisements[i].AggregationLengthV6 = *adv.AggregationLengthV6
		dst.Spec.BGPAdvertisements[i].LocalPref = adv.LocalPref
		dst.Spec.BGPAdvertisements[i].Communities = make([]string, len(adv.Communities))
		copy(dst.Spec.BGPAdvertisements[i].Communities, adv.Communities)
	}
	dst.ObjectMeta = src.ObjectMeta
	return nil
}

// ConvertFrom converts from the Hub version (vbeta1) to this version.
func (dst *AddressPool) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1beta1.AddressPool)
	dst.ObjectMeta = src.ObjectMeta

	dst.Spec.Protocol = src.Spec.Protocol
	dst.Spec.Addresses = make([]string, len(src.Spec.Addresses))
	copy(dst.Spec.Addresses, src.Spec.Addresses)
	*dst.Spec.AutoAssign = *src.Spec.AutoAssign
	dst.Spec.BGPAdvertisements = make([]BgpAdvertisement, 0)
	for i, adv := range src.Spec.BGPAdvertisements {
		*dst.Spec.BGPAdvertisements[i].AggregationLength = *adv.AggregationLength
		*dst.Spec.BGPAdvertisements[i].AggregationLengthV6 = *adv.AggregationLengthV6
		dst.Spec.BGPAdvertisements[i].LocalPref = adv.LocalPref
		dst.Spec.BGPAdvertisements[i].Communities = make([]string, len(adv.Communities))
		copy(dst.Spec.BGPAdvertisements[i].Communities, adv.Communities)
	}
	return nil
}
