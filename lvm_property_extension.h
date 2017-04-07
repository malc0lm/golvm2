#include <lvm2app.h>


typedef struct lvm_property_value_c {
	uint32_t is_settable;
	uint32_t is_string;
	uint32_t is_integer;
	uint32_t is_valid;
	union {
		const char *c_string;
		uint64_t integer;
	} value;
} lvm_property_value_c;

lvm_property_value_c *lvm_vg_get_property_c(const vg_t vg, const char *name);
lvm_property_value_c *lvm_lvseg_get_property_c(const lvseg_t lvseg, const char *name);
lvm_property_value_c *lvm_lv_get_property_c(const lv_t lv, const char *name);
lvm_property_value_c *lvm_pv_get_property_c(const pv_t pv, const char *name);
lvm_property_value_c *lvm_pvseg_get_property_c(const pvseg_t pvseg, const char *name);