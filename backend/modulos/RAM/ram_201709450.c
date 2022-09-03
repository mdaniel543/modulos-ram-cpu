#include <linux/module.h>
// para usar KERN_INFO
#include <linux/kernel.h>

//Header para los macros module_init y module_exit
#include <linux/init.h>
//Header necesario porque se usara proc_fs
#include <linux/proc_fs.h>
/* for copy_from_user */
#include <asm/uaccess.h>	
/* Header para usar la lib seq_file y manejar el archivo en /proc*/
#include <linux/seq_file.h>

#include <linux/mm.h>

MODULE_LICENSE("GPL");
MODULE_DESCRIPTION("Practica 2 Sistemas Operativos 1"); 
MODULE_AUTHOR("Marvin Daniel Rodriguez Felix"); 

//Funcion que se ejecutara cada vez que se lea el archivo con el comando CAT
static int escribir_archivo(struct seq_file *file, void *v)
{   
    struct sysinfo inf;
    si_meminfo(&inf);
    seq_printf(file, "{\n");
    seq_printf(file,"\"total\": %lu,\n",inf.totalram*4/1024);
    seq_printf(file,"\"free\": %lu,\n", inf.freeram*4/1024);
    seq_printf(file, "\"used\": %lu,\n", (inf.totalram - inf.freeram)* 4/1024); 
    seq_printf(file,"\"percentage\": %lu\n", ((inf.totalram - inf.freeram)*100)/inf.totalram);
    seq_printf(file, "}\n");
    return 0;
}

//Funcion que se ejecutara cada vez que se lea el archivo con el comando CAT
static int al_abrir(struct inode *inode, struct file *file)
{
    return single_open(file, escribir_archivo, NULL);
}

//Si el kernel es 5.6 o mayor se usa la estructura proc_ops
static struct proc_ops operaciones =
{
    .proc_open = al_abrir,
    .proc_read = seq_read
};

//Funcion a ejecuta al insertar el modulo en el kernel con insmod
static int _insert(void)
{
    proc_create("ram_201709450", 0, NULL, &operaciones);
    printk(KERN_INFO "Sistemas Operativos 1\n");
    return 0;
}

//Funcion a ejecuta al remover el modulo del kernel con rmmod
static void _remove(void)
{
    remove_proc_entry("ram_201709450", NULL);
    printk(KERN_INFO "20170950\n");
}

module_init(_insert);
module_exit(_remove);