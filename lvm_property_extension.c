#include <lvm2app.h>
#include "lvm_property_extension.h"

/*
   We must extend the lvm2app.h because Cgo didn't support C struct bit fields.
   And we have to release memory which allocate by malloc in C manually.
*/

lvm_property_value_c *lvm_vg_get_property_c(const vg_t vg, const char *name){
	lvm_property_value_t v;
	v = lvm_vg_get_property(vg, name);
	if (!v.is_valid) {
		return NULL;
	}
	lvm_property_value_c *prop = (lvm_property_value_c *)malloc(sizeof(lvm_property_value_c));

	prop->is_integer = v.is_integer;
	prop->is_settable = v.is_settable;
	prop->is_string = v.is_string;
	prop->is_valid = v.is_valid;
	if (v.is_string){
		prop->value.c_string = v.value.string;
	}
	if (v.is_integer){
		prop->value.integer = v.value.integer;
	}
	return prop;
}

lvm_property_value_c *lvm_lvseg_get_property_c(const lvseg_t lvseg, const char *name){
	lvm_property_value_t v;
	v = lvm_lvseg_get_property(lvseg, name);
	if (!v.is_valid) {
		return NULL;
	}
	lvm_property_value_c *prop = (lvm_property_value_c *)malloc(sizeof(lvm_property_value_c));

	prop->is_integer = v.is_integer;
	prop->is_settable = v.is_settable;
	prop->is_string = v.is_string;
	prop->is_valid = v.is_valid;
	if (v.is_string){
		prop->value.c_string = v.value.string;
	}
	if (v.is_integer){
		prop->value.integer = v.value.integer;
	}
	return prop;
}

lvm_property_value_c *lvm_lv_get_property_c(const lv_t lv, const char *name){
	lvm_property_value_t v;
	v = lvm_lv_get_property(lv, name);
	if (!v.is_valid) {
		return NULL;
	}
	lvm_property_value_c *prop = (lvm_property_value_c *)malloc(sizeof(lvm_property_value_c));

	prop->is_integer = v.is_integer;
	prop->is_settable = v.is_settable;
	prop->is_string = v.is_string;
	prop->is_valid = v.is_valid;
	if (v.is_string){
		prop->value.c_string = v.value.string;
	}
	if (v.is_integer){
		prop->value.integer = v.value.integer;
	}
	return prop;
}

lvm_property_value_c *lvm_pv_get_property_c(const pv_t pv, const char *name){
	lvm_property_value_t v;
	v = lvm_pv_get_property(pv, name);
	if (!v.is_valid) {
		return NULL;
	}
	lvm_property_value_c *prop = (lvm_property_value_c *)malloc(sizeof(lvm_property_value_c));

	prop->is_integer = v.is_integer;
	prop->is_settable = v.is_settable;
	prop->is_string = v.is_string;
	prop->is_valid = v.is_valid;
	if (v.is_string){
		prop->value.c_string = v.value.string;
	}
	if (v.is_integer){
		prop->value.integer = v.value.integer;
	}
	return prop;
}

lvm_property_value_c *lvm_pvseg_get_property_c(const pvseg_t pvseg, const char *name){
	lvm_property_value_t v;
	v = lvm_pvseg_get_property(pvseg, name);
	if (!v.is_valid) {
		return NULL;
	}
	lvm_property_value_c *prop = (lvm_property_value_c *)malloc(sizeof(lvm_property_value_c));

	prop->is_integer = v.is_integer;
	prop->is_settable = v.is_settable;
	prop->is_string = v.is_string;
	prop->is_valid = v.is_valid;
	if (v.is_string){
		prop->value.c_string = v.value.string;
	}
	if (v.is_integer){
		prop->value.integer = v.value.integer;
	}
	return prop;
}
