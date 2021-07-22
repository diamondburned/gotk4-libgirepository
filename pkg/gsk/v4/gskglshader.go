// Code generated by girgen. DO NOT EDIT.

package gsk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/graphene"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gsk/gsk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gsk_gl_shader_get_type()), F: marshalGLShaderer},
		{T: externglib.Type(C.gsk_shader_args_builder_get_type()), F: marshalShaderArgsBuilder},
	})
}

// GLShader: GskGLShader is a snippet of GLSL that is meant to run in the
// fragment shader of the rendering pipeline.
//
// A fragment shader gets the coordinates being rendered as input and produces
// the pixel values for that particular pixel. Additionally, the shader can
// declare a set of other input arguments, called uniforms (as they are uniform
// over all the calls to your shader in each instance of use). A shader can also
// receive up to 4 textures that it can use as input when producing the pixel
// data.
//
// GskGLShader is usually used with gtk_snapshot_push_gl_shader() to produce a
// gsk.GLShaderNode in the rendering hierarchy, and then its input textures are
// constructed by rendering the child nodes to textures before rendering the
// shader node itself. (You can pass texture nodes as children if you want to
// directly use a texture as input).
//
// The actual shader code is GLSL code that gets combined with some other code
// into the fragment shader. Since the exact capabilities of the GPU driver
// differs between different OpenGL drivers and hardware, GTK adds some defines
// that you can use to ensure your GLSL code runs on as many drivers as it can.
//
// If the OpenGL driver is GLES, then the shader language version is set to 100,
// and GSK_GLES will be defined in the shader.
//
// Otherwise, if the OpenGL driver does not support the 3.2 core profile, then
// the shader will run with language version 110 for GL2 and 130 for GL3, and
// GSK_LEGACY will be defined in the shader.
//
// If the OpenGL driver supports the 3.2 code profile, it will be used, the
// shader language version is set to 150, and GSK_GL3 will be defined in the
// shader.
//
// The main function the shader must implement is:
//
//     void mainImage(out vec4 fragColor,
//                    in vec2 fragCoord,
//                    in vec2 resolution,
//                    in vec2 uv)
//
//
// Where the input fragCoord is the coordinate of the pixel we're currently
// rendering, relative to the boundary rectangle that was specified in the
// GskGLShaderNode, and resolution is the width and height of that rectangle.
// This is in the typical GTK coordinate system with the origin in the top left.
// uv contains the u and v coordinates that can be used to index a texture at
// the corresponding point. These coordinates are in the [0..1]x[0..1] region,
// with 0, 0 being in the lower left corder (which is typical for OpenGL).
//
// The output fragColor should be a RGBA color (with premultiplied alpha) that
// will be used as the output for the specified pixel location. Note that this
// output will be automatically clipped to the clip region of the glshader node.
//
// In addition to the function arguments the shader can define up to 4 uniforms
// for textures which must be called u_textureN (i.e. u_texture1 to u_texture4)
// as well as any custom uniforms you want of types int, uint, bool, float,
// vec2, vec3 or vec4.
//
// All textures sources contain premultiplied alpha colors, but if some there
// are outer sources of colors there is a gsk_premultiply() helper to compute
// premultiplication when needed.
//
// Note that GTK parses the uniform declarations, so each uniform has to be on a
// line by itself with no other code, like so:
//
//    uniform float u_time;
//    uniform vec3 u_color;
//    uniform sampler2D u_texture1;
//    uniform sampler2D u_texture2;
//
//
// GTK uses the the "gsk" namespace in the symbols it uses in the shader, so
// your code should not use any symbols with the prefix gsk or GSK. There are
// some helper functions declared that you can use:
//
//    vec4 GskTexture(sampler2D sampler, vec2 texCoords);
//
//
// This samples a texture (e.g. u_texture1) at the specified coordinates, and
// containes some helper ifdefs to ensure that it works on all OpenGL versions.
//
// You can compile the shader yourself using gsk.GLShader.Compile(), otherwise
// the GSK renderer will do it when it handling the glshader node. If errors
// occurs, the returned error will include the glsl sources, so you can see what
// GSK was passing to the compiler. You can also set GSK_DEBUG=shaders in the
// environment to see the sources and other relevant information about all
// shaders that GSK is handling.
//
// An example shader
//
//    uniform float position;
//    uniform sampler2D u_texture1;
//    uniform sampler2D u_texture2;
//
//    void mainImage(out vec4 fragColor,
//                   in vec2 fragCoord,
//                   in vec2 resolution,
//                   in vec2 uv) {
//      vec4 source1 = GskTexture(u_texture1, uv);
//      vec4 source2 = GskTexture(u_texture2, uv);
//
//      fragColor = position * source1 + (1.0 - position) * source2;
//    }
type GLShader struct {
	*externglib.Object
}

func wrapGLShader(obj *externglib.Object) *GLShader {
	return &GLShader{
		Object: obj,
	}
}

func marshalGLShaderer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapGLShader(obj), nil
}

// NewGLShaderFromResource creates a GskGLShader that will render pixels using
// the specified code.
func NewGLShaderFromResource(resourcePath string) *GLShader {
	var _arg1 *C.char        // out
	var _cret *C.GskGLShader // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(resourcePath)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gsk_gl_shader_new_from_resource(_arg1)

	var _glShader *GLShader // out

	_glShader = wrapGLShader(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _glShader
}

// Compile tries to compile the shader for the given renderer.
//
// If there is a problem, this function returns FALSE and reports an error. You
// should use this function before relying on the shader for rendering and use a
// fallback with a simpler shader or without shaders if it fails.
//
// Note that this will modify the rendering state (for example change the
// current GL context) and requires the renderer to be set up. This means that
// the widget has to be realized. Commonly you want to call this from the
// realize signal of a widget, or during widget snapshot.
func (shader *GLShader) Compile(renderer Rendererer) error {
	var _arg0 *C.GskGLShader // out
	var _arg1 *C.GskRenderer // out
	var _cerr *C.GError      // in

	_arg0 = (*C.GskGLShader)(unsafe.Pointer(shader.Native()))
	_arg1 = (*C.GskRenderer)(unsafe.Pointer(renderer.Native()))

	C.gsk_gl_shader_compile(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// FindUniformByName looks for a uniform by the name name, and returns the index
// of the uniform, or -1 if it was not found.
func (shader *GLShader) FindUniformByName(name string) int {
	var _arg0 *C.GskGLShader // out
	var _arg1 *C.char        // out
	var _cret C.int          // in

	_arg0 = (*C.GskGLShader)(unsafe.Pointer(shader.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gsk_gl_shader_find_uniform_by_name(_arg0, _arg1)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// ArgsSize: get the size of the data block used to specify arguments for this
// shader.
func (shader *GLShader) ArgsSize() uint {
	var _arg0 *C.GskGLShader // out
	var _cret C.gsize        // in

	_arg0 = (*C.GskGLShader)(unsafe.Pointer(shader.Native()))

	_cret = C.gsk_gl_shader_get_args_size(_arg0)

	var _gsize uint // out

	_gsize = uint(_cret)

	return _gsize
}

// NTextures returns the number of textures that the shader requires.
//
// This can be used to check that the a passed shader works in your usecase. It
// is determined by looking at the highest u_textureN value that the shader
// defines.
func (shader *GLShader) NTextures() int {
	var _arg0 *C.GskGLShader // out
	var _cret C.int          // in

	_arg0 = (*C.GskGLShader)(unsafe.Pointer(shader.Native()))

	_cret = C.gsk_gl_shader_get_n_textures(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// NUniforms: get the number of declared uniforms for this shader.
func (shader *GLShader) NUniforms() int {
	var _arg0 *C.GskGLShader // out
	var _cret C.int          // in

	_arg0 = (*C.GskGLShader)(unsafe.Pointer(shader.Native()))

	_cret = C.gsk_gl_shader_get_n_uniforms(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Resource gets the resource path for the GLSL sourcecode being used to render
// this shader.
func (shader *GLShader) Resource() string {
	var _arg0 *C.GskGLShader // out
	var _cret *C.char        // in

	_arg0 = (*C.GskGLShader)(unsafe.Pointer(shader.Native()))

	_cret = C.gsk_gl_shader_get_resource(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// UniformName: get the name of the declared uniform for this shader at index
// idx.
func (shader *GLShader) UniformName(idx int) string {
	var _arg0 *C.GskGLShader // out
	var _arg1 C.int          // out
	var _cret *C.char        // in

	_arg0 = (*C.GskGLShader)(unsafe.Pointer(shader.Native()))
	_arg1 = C.int(idx)

	_cret = C.gsk_gl_shader_get_uniform_name(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// UniformOffset: get the offset into the data block where data for this
// uniforms is stored.
func (shader *GLShader) UniformOffset(idx int) int {
	var _arg0 *C.GskGLShader // out
	var _arg1 C.int          // out
	var _cret C.int          // in

	_arg0 = (*C.GskGLShader)(unsafe.Pointer(shader.Native()))
	_arg1 = C.int(idx)

	_cret = C.gsk_gl_shader_get_uniform_offset(_arg0, _arg1)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// UniformType: get the type of the declared uniform for this shader at index
// idx.
func (shader *GLShader) UniformType(idx int) GLUniformType {
	var _arg0 *C.GskGLShader     // out
	var _arg1 C.int              // out
	var _cret C.GskGLUniformType // in

	_arg0 = (*C.GskGLShader)(unsafe.Pointer(shader.Native()))
	_arg1 = C.int(idx)

	_cret = C.gsk_gl_shader_get_uniform_type(_arg0, _arg1)

	var _glUniformType GLUniformType // out

	_glUniformType = GLUniformType(_cret)

	return _glUniformType
}

// ShaderArgsBuilder: object to build the uniforms data for a GLShader.
type ShaderArgsBuilder struct {
	nocopy gextras.NoCopy
	native *C.GskShaderArgsBuilder
}

func marshalShaderArgsBuilder(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return &ShaderArgsBuilder{native: (*C.GskShaderArgsBuilder)(unsafe.Pointer(b))}, nil
}

// SetBool sets the value of the uniform idx.
//
// The uniform must be of bool type.
func (builder *ShaderArgsBuilder) SetBool(idx int, value bool) {
	var _arg0 *C.GskShaderArgsBuilder // out
	var _arg1 C.int                   // out
	var _arg2 C.gboolean              // out

	_arg0 = (*C.GskShaderArgsBuilder)(gextras.StructNative(unsafe.Pointer(builder)))
	_arg1 = C.int(idx)
	if value {
		_arg2 = C.TRUE
	}

	C.gsk_shader_args_builder_set_bool(_arg0, _arg1, _arg2)
}

// SetFloat sets the value of the uniform idx.
//
// The uniform must be of float type.
func (builder *ShaderArgsBuilder) SetFloat(idx int, value float32) {
	var _arg0 *C.GskShaderArgsBuilder // out
	var _arg1 C.int                   // out
	var _arg2 C.float                 // out

	_arg0 = (*C.GskShaderArgsBuilder)(gextras.StructNative(unsafe.Pointer(builder)))
	_arg1 = C.int(idx)
	_arg2 = C.float(value)

	C.gsk_shader_args_builder_set_float(_arg0, _arg1, _arg2)
}

// SetInt sets the value of the uniform idx.
//
// The uniform must be of int type.
func (builder *ShaderArgsBuilder) SetInt(idx int, value int32) {
	var _arg0 *C.GskShaderArgsBuilder // out
	var _arg1 C.int                   // out
	var _arg2 C.gint32                // out

	_arg0 = (*C.GskShaderArgsBuilder)(gextras.StructNative(unsafe.Pointer(builder)))
	_arg1 = C.int(idx)
	_arg2 = C.gint32(value)

	C.gsk_shader_args_builder_set_int(_arg0, _arg1, _arg2)
}

// SetUint sets the value of the uniform idx.
//
// The uniform must be of uint type.
func (builder *ShaderArgsBuilder) SetUint(idx int, value uint32) {
	var _arg0 *C.GskShaderArgsBuilder // out
	var _arg1 C.int                   // out
	var _arg2 C.guint32               // out

	_arg0 = (*C.GskShaderArgsBuilder)(gextras.StructNative(unsafe.Pointer(builder)))
	_arg1 = C.int(idx)
	_arg2 = C.guint32(value)

	C.gsk_shader_args_builder_set_uint(_arg0, _arg1, _arg2)
}

// SetVec2 sets the value of the uniform idx.
//
// The uniform must be of vec2 type.
func (builder *ShaderArgsBuilder) SetVec2(idx int, value *graphene.Vec2) {
	var _arg0 *C.GskShaderArgsBuilder // out
	var _arg1 C.int                   // out
	var _arg2 *C.graphene_vec2_t      // out

	_arg0 = (*C.GskShaderArgsBuilder)(gextras.StructNative(unsafe.Pointer(builder)))
	_arg1 = C.int(idx)
	_arg2 = (*C.graphene_vec2_t)(gextras.StructNative(unsafe.Pointer(value)))

	C.gsk_shader_args_builder_set_vec2(_arg0, _arg1, _arg2)
}

// SetVec3 sets the value of the uniform idx.
//
// The uniform must be of vec3 type.
func (builder *ShaderArgsBuilder) SetVec3(idx int, value *graphene.Vec3) {
	var _arg0 *C.GskShaderArgsBuilder // out
	var _arg1 C.int                   // out
	var _arg2 *C.graphene_vec3_t      // out

	_arg0 = (*C.GskShaderArgsBuilder)(gextras.StructNative(unsafe.Pointer(builder)))
	_arg1 = C.int(idx)
	_arg2 = (*C.graphene_vec3_t)(gextras.StructNative(unsafe.Pointer(value)))

	C.gsk_shader_args_builder_set_vec3(_arg0, _arg1, _arg2)
}

// SetVec4 sets the value of the uniform idx.
//
// The uniform must be of vec4 type.
func (builder *ShaderArgsBuilder) SetVec4(idx int, value *graphene.Vec4) {
	var _arg0 *C.GskShaderArgsBuilder // out
	var _arg1 C.int                   // out
	var _arg2 *C.graphene_vec4_t      // out

	_arg0 = (*C.GskShaderArgsBuilder)(gextras.StructNative(unsafe.Pointer(builder)))
	_arg1 = C.int(idx)
	_arg2 = (*C.graphene_vec4_t)(gextras.StructNative(unsafe.Pointer(value)))

	C.gsk_shader_args_builder_set_vec4(_arg0, _arg1, _arg2)
}
