// Code generated by girgen. DO NOT EDIT.

package gsk

import (
	"sync"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gsk/gsk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gsk_blend_node_get_type()), F: marshalBlendNoder},
		{T: externglib.Type(C.gsk_blur_node_get_type()), F: marshalBlurNoder},
		{T: externglib.Type(C.gsk_border_node_get_type()), F: marshalBorderNoder},
		{T: externglib.Type(C.gsk_cairo_node_get_type()), F: marshalCairoNoder},
		{T: externglib.Type(C.gsk_clip_node_get_type()), F: marshalClipNoder},
		{T: externglib.Type(C.gsk_color_matrix_node_get_type()), F: marshalColorMatrixNoder},
		{T: externglib.Type(C.gsk_color_node_get_type()), F: marshalColorNoder},
		{T: externglib.Type(C.gsk_conic_gradient_node_get_type()), F: marshalConicGradientNoder},
		{T: externglib.Type(C.gsk_container_node_get_type()), F: marshalContainerNoder},
		{T: externglib.Type(C.gsk_cross_fade_node_get_type()), F: marshalCrossFadeNoder},
		{T: externglib.Type(C.gsk_debug_node_get_type()), F: marshalDebugNoder},
		{T: externglib.Type(C.gsk_gl_shader_node_get_type()), F: marshalGLShaderNoder},
		{T: externglib.Type(C.gsk_inset_shadow_node_get_type()), F: marshalInsetShadowNoder},
		{T: externglib.Type(C.gsk_linear_gradient_node_get_type()), F: marshalLinearGradientNoder},
		{T: externglib.Type(C.gsk_opacity_node_get_type()), F: marshalOpacityNoder},
		{T: externglib.Type(C.gsk_outset_shadow_node_get_type()), F: marshalOutsetShadowNoder},
		{T: externglib.Type(C.gsk_radial_gradient_node_get_type()), F: marshalRadialGradientNoder},
		{T: externglib.Type(C.gsk_repeat_node_get_type()), F: marshalRepeatNoder},
		{T: externglib.Type(C.gsk_repeating_linear_gradient_node_get_type()), F: marshalRepeatingLinearGradientNoder},
		{T: externglib.Type(C.gsk_repeating_radial_gradient_node_get_type()), F: marshalRepeatingRadialGradientNoder},
		{T: externglib.Type(C.gsk_rounded_clip_node_get_type()), F: marshalRoundedClipNoder},
		{T: externglib.Type(C.gsk_shadow_node_get_type()), F: marshalShadowNoder},
		{T: externglib.Type(C.gsk_text_node_get_type()), F: marshalTextNoder},
		{T: externglib.Type(C.gsk_texture_node_get_type()), F: marshalTextureNoder},
		{T: externglib.Type(C.gsk_transform_node_get_type()), F: marshalTransformNoder},
	})
}

// BlendNode: render node applying a blending function between its two child
// nodes.
type BlendNode struct {
	RenderNode

	_ [0]func()     // equal guard
	_ [0]sync.Mutex // copy guard
}

var (
	_ RenderNoder = (*BlendNode)(nil)
)

func wrapBlendNode(obj *externglib.Object) *BlendNode {
	return &BlendNode{
		RenderNode: RenderNode{
			Object: obj,
		},
	}
}

func marshalBlendNoder(p uintptr) (interface{}, error) {
	return wrapBlendNode(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// BlurNode: render node applying a blur effect to its single child.
type BlurNode struct {
	RenderNode

	_ [0]func()     // equal guard
	_ [0]sync.Mutex // copy guard
}

var (
	_ RenderNoder = (*BlurNode)(nil)
)

func wrapBlurNode(obj *externglib.Object) *BlurNode {
	return &BlurNode{
		RenderNode: RenderNode{
			Object: obj,
		},
	}
}

func marshalBlurNoder(p uintptr) (interface{}, error) {
	return wrapBlurNode(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// BorderNode: render node for a border.
type BorderNode struct {
	RenderNode

	_ [0]func()     // equal guard
	_ [0]sync.Mutex // copy guard
}

var (
	_ RenderNoder = (*BorderNode)(nil)
)

func wrapBorderNode(obj *externglib.Object) *BorderNode {
	return &BorderNode{
		RenderNode: RenderNode{
			Object: obj,
		},
	}
}

func marshalBorderNoder(p uintptr) (interface{}, error) {
	return wrapBorderNode(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// CairoNode: render node for a Cairo surface.
type CairoNode struct {
	RenderNode

	_ [0]func()     // equal guard
	_ [0]sync.Mutex // copy guard
}

var (
	_ RenderNoder = (*CairoNode)(nil)
)

func wrapCairoNode(obj *externglib.Object) *CairoNode {
	return &CairoNode{
		RenderNode: RenderNode{
			Object: obj,
		},
	}
}

func marshalCairoNoder(p uintptr) (interface{}, error) {
	return wrapCairoNode(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ClipNode: render node applying a rectangular clip to its single child node.
type ClipNode struct {
	RenderNode

	_ [0]func()     // equal guard
	_ [0]sync.Mutex // copy guard
}

var (
	_ RenderNoder = (*ClipNode)(nil)
)

func wrapClipNode(obj *externglib.Object) *ClipNode {
	return &ClipNode{
		RenderNode: RenderNode{
			Object: obj,
		},
	}
}

func marshalClipNoder(p uintptr) (interface{}, error) {
	return wrapClipNode(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ColorMatrixNode: render node controlling the color matrix of its single child
// node.
type ColorMatrixNode struct {
	RenderNode

	_ [0]func()     // equal guard
	_ [0]sync.Mutex // copy guard
}

var (
	_ RenderNoder = (*ColorMatrixNode)(nil)
)

func wrapColorMatrixNode(obj *externglib.Object) *ColorMatrixNode {
	return &ColorMatrixNode{
		RenderNode: RenderNode{
			Object: obj,
		},
	}
}

func marshalColorMatrixNoder(p uintptr) (interface{}, error) {
	return wrapColorMatrixNode(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ColorNode: render node for a solid color.
type ColorNode struct {
	RenderNode

	_ [0]func()     // equal guard
	_ [0]sync.Mutex // copy guard
}

var (
	_ RenderNoder = (*ColorNode)(nil)
)

func wrapColorNode(obj *externglib.Object) *ColorNode {
	return &ColorNode{
		RenderNode: RenderNode{
			Object: obj,
		},
	}
}

func marshalColorNoder(p uintptr) (interface{}, error) {
	return wrapColorNode(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConicGradientNode: render node for a conic gradient.
type ConicGradientNode struct {
	RenderNode

	_ [0]func()     // equal guard
	_ [0]sync.Mutex // copy guard
}

var (
	_ RenderNoder = (*ConicGradientNode)(nil)
)

func wrapConicGradientNode(obj *externglib.Object) *ConicGradientNode {
	return &ConicGradientNode{
		RenderNode: RenderNode{
			Object: obj,
		},
	}
}

func marshalConicGradientNoder(p uintptr) (interface{}, error) {
	return wrapConicGradientNode(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ContainerNode: render node that can contain other render nodes.
type ContainerNode struct {
	RenderNode

	_ [0]func()     // equal guard
	_ [0]sync.Mutex // copy guard
}

var (
	_ RenderNoder = (*ContainerNode)(nil)
)

func wrapContainerNode(obj *externglib.Object) *ContainerNode {
	return &ContainerNode{
		RenderNode: RenderNode{
			Object: obj,
		},
	}
}

func marshalContainerNoder(p uintptr) (interface{}, error) {
	return wrapContainerNode(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// CrossFadeNode: render node cross fading between two child nodes.
type CrossFadeNode struct {
	RenderNode

	_ [0]func()     // equal guard
	_ [0]sync.Mutex // copy guard
}

var (
	_ RenderNoder = (*CrossFadeNode)(nil)
)

func wrapCrossFadeNode(obj *externglib.Object) *CrossFadeNode {
	return &CrossFadeNode{
		RenderNode: RenderNode{
			Object: obj,
		},
	}
}

func marshalCrossFadeNoder(p uintptr) (interface{}, error) {
	return wrapCrossFadeNode(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// DebugNode: render node that emits a debugging message when drawing its child
// node.
type DebugNode struct {
	RenderNode

	_ [0]func()     // equal guard
	_ [0]sync.Mutex // copy guard
}

var (
	_ RenderNoder = (*DebugNode)(nil)
)

func wrapDebugNode(obj *externglib.Object) *DebugNode {
	return &DebugNode{
		RenderNode: RenderNode{
			Object: obj,
		},
	}
}

func marshalDebugNoder(p uintptr) (interface{}, error) {
	return wrapDebugNode(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// GLShaderNode: render node using a GL shader when drawing its children nodes.
type GLShaderNode struct {
	RenderNode

	_ [0]func()     // equal guard
	_ [0]sync.Mutex // copy guard
}

var (
	_ RenderNoder = (*GLShaderNode)(nil)
)

func wrapGLShaderNode(obj *externglib.Object) *GLShaderNode {
	return &GLShaderNode{
		RenderNode: RenderNode{
			Object: obj,
		},
	}
}

func marshalGLShaderNoder(p uintptr) (interface{}, error) {
	return wrapGLShaderNode(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// InsetShadowNode: render node for an inset shadow.
type InsetShadowNode struct {
	RenderNode

	_ [0]func()     // equal guard
	_ [0]sync.Mutex // copy guard
}

var (
	_ RenderNoder = (*InsetShadowNode)(nil)
)

func wrapInsetShadowNode(obj *externglib.Object) *InsetShadowNode {
	return &InsetShadowNode{
		RenderNode: RenderNode{
			Object: obj,
		},
	}
}

func marshalInsetShadowNoder(p uintptr) (interface{}, error) {
	return wrapInsetShadowNode(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// LinearGradientNode: render node for a linear gradient.
type LinearGradientNode struct {
	RenderNode

	_ [0]func()     // equal guard
	_ [0]sync.Mutex // copy guard
}

var (
	_ RenderNoder = (*LinearGradientNode)(nil)
)

func wrapLinearGradientNode(obj *externglib.Object) *LinearGradientNode {
	return &LinearGradientNode{
		RenderNode: RenderNode{
			Object: obj,
		},
	}
}

func marshalLinearGradientNoder(p uintptr) (interface{}, error) {
	return wrapLinearGradientNode(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// OpacityNode: render node controlling the opacity of its single child node.
type OpacityNode struct {
	RenderNode

	_ [0]func()     // equal guard
	_ [0]sync.Mutex // copy guard
}

var (
	_ RenderNoder = (*OpacityNode)(nil)
)

func wrapOpacityNode(obj *externglib.Object) *OpacityNode {
	return &OpacityNode{
		RenderNode: RenderNode{
			Object: obj,
		},
	}
}

func marshalOpacityNoder(p uintptr) (interface{}, error) {
	return wrapOpacityNode(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// OutsetShadowNode: render node for an outset shadow.
type OutsetShadowNode struct {
	RenderNode

	_ [0]func()     // equal guard
	_ [0]sync.Mutex // copy guard
}

var (
	_ RenderNoder = (*OutsetShadowNode)(nil)
)

func wrapOutsetShadowNode(obj *externglib.Object) *OutsetShadowNode {
	return &OutsetShadowNode{
		RenderNode: RenderNode{
			Object: obj,
		},
	}
}

func marshalOutsetShadowNoder(p uintptr) (interface{}, error) {
	return wrapOutsetShadowNode(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// RadialGradientNode: render node for a radial gradient.
type RadialGradientNode struct {
	RenderNode

	_ [0]func()     // equal guard
	_ [0]sync.Mutex // copy guard
}

var (
	_ RenderNoder = (*RadialGradientNode)(nil)
)

func wrapRadialGradientNode(obj *externglib.Object) *RadialGradientNode {
	return &RadialGradientNode{
		RenderNode: RenderNode{
			Object: obj,
		},
	}
}

func marshalRadialGradientNoder(p uintptr) (interface{}, error) {
	return wrapRadialGradientNode(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// RepeatNode: render node repeating its single child node.
type RepeatNode struct {
	RenderNode

	_ [0]func()     // equal guard
	_ [0]sync.Mutex // copy guard
}

var (
	_ RenderNoder = (*RepeatNode)(nil)
)

func wrapRepeatNode(obj *externglib.Object) *RepeatNode {
	return &RepeatNode{
		RenderNode: RenderNode{
			Object: obj,
		},
	}
}

func marshalRepeatNoder(p uintptr) (interface{}, error) {
	return wrapRepeatNode(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// RepeatingLinearGradientNode: render node for a repeating linear gradient.
type RepeatingLinearGradientNode struct {
	RenderNode

	_ [0]func()     // equal guard
	_ [0]sync.Mutex // copy guard
}

var (
	_ RenderNoder = (*RepeatingLinearGradientNode)(nil)
)

func wrapRepeatingLinearGradientNode(obj *externglib.Object) *RepeatingLinearGradientNode {
	return &RepeatingLinearGradientNode{
		RenderNode: RenderNode{
			Object: obj,
		},
	}
}

func marshalRepeatingLinearGradientNoder(p uintptr) (interface{}, error) {
	return wrapRepeatingLinearGradientNode(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// RepeatingRadialGradientNode: render node for a repeating radial gradient.
type RepeatingRadialGradientNode struct {
	RenderNode

	_ [0]func()     // equal guard
	_ [0]sync.Mutex // copy guard
}

var (
	_ RenderNoder = (*RepeatingRadialGradientNode)(nil)
)

func wrapRepeatingRadialGradientNode(obj *externglib.Object) *RepeatingRadialGradientNode {
	return &RepeatingRadialGradientNode{
		RenderNode: RenderNode{
			Object: obj,
		},
	}
}

func marshalRepeatingRadialGradientNoder(p uintptr) (interface{}, error) {
	return wrapRepeatingRadialGradientNode(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// RoundedClipNode: render node applying a rounded rectangle clip to its single
// child.
type RoundedClipNode struct {
	RenderNode

	_ [0]func()     // equal guard
	_ [0]sync.Mutex // copy guard
}

var (
	_ RenderNoder = (*RoundedClipNode)(nil)
)

func wrapRoundedClipNode(obj *externglib.Object) *RoundedClipNode {
	return &RoundedClipNode{
		RenderNode: RenderNode{
			Object: obj,
		},
	}
}

func marshalRoundedClipNoder(p uintptr) (interface{}, error) {
	return wrapRoundedClipNode(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ShadowNode: render node drawing one or more shadows behind its single child
// node.
type ShadowNode struct {
	RenderNode

	_ [0]func()     // equal guard
	_ [0]sync.Mutex // copy guard
}

var (
	_ RenderNoder = (*ShadowNode)(nil)
)

func wrapShadowNode(obj *externglib.Object) *ShadowNode {
	return &ShadowNode{
		RenderNode: RenderNode{
			Object: obj,
		},
	}
}

func marshalShadowNoder(p uintptr) (interface{}, error) {
	return wrapShadowNode(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// TextNode: render node drawing a set of glyphs.
type TextNode struct {
	RenderNode

	_ [0]func()     // equal guard
	_ [0]sync.Mutex // copy guard
}

var (
	_ RenderNoder = (*TextNode)(nil)
)

func wrapTextNode(obj *externglib.Object) *TextNode {
	return &TextNode{
		RenderNode: RenderNode{
			Object: obj,
		},
	}
}

func marshalTextNoder(p uintptr) (interface{}, error) {
	return wrapTextNode(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// TextureNode: render node for a Texture.
type TextureNode struct {
	RenderNode

	_ [0]func()     // equal guard
	_ [0]sync.Mutex // copy guard
}

var (
	_ RenderNoder = (*TextureNode)(nil)
)

func wrapTextureNode(obj *externglib.Object) *TextureNode {
	return &TextureNode{
		RenderNode: RenderNode{
			Object: obj,
		},
	}
}

func marshalTextureNoder(p uintptr) (interface{}, error) {
	return wrapTextureNode(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// TransformNode: render node applying a GskTransform to its single child node.
type TransformNode struct {
	RenderNode

	_ [0]func()     // equal guard
	_ [0]sync.Mutex // copy guard
}

var (
	_ RenderNoder = (*TransformNode)(nil)
)

func wrapTransformNode(obj *externglib.Object) *TransformNode {
	return &TransformNode{
		RenderNode: RenderNode{
			Object: obj,
		},
	}
}

func marshalTransformNoder(p uintptr) (interface{}, error) {
	return wrapTransformNode(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
