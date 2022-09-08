#include <linux/init.h>
#include <linux/module.h>
#include <linux/proc_fs.h>
#include <linux/seq_file.h>
#include <linux/sched/signal.h>
#include <linux/sched.h>
#include <linux/fs.h>
#include <linux/sched/mm.h> // get_task_mm(), mmput()
#include <linux/mm.h>       // get_mm_rss()

MODULE_LICENSE("GPL");
MODULE_DESCRIPTION("Practica 2 Modulo CPU");
MODULE_AUTHOR("Marvin Daniel Rodriguez Felix");

struct task_struct *cpu;
struct task_struct *child;
struct list_head *lstProcess;
struct mm_struct *mm;

// Funcion que se ejecutara cada vez que se lea el archivo con el comando CAT
static int escribir_archivo(struct seq_file *archivo, void *v)
{
    seq_printf(archivo, "[\n");
    for_each_process(cpu)
    {
        seq_printf(archivo, "{ \n");
        seq_printf(archivo, "\"pid\":%d,\n", cpu->pid);
        seq_printf(archivo, "\"name\":\"%s\",\n", cpu->comm);
        seq_printf(archivo, "\"user\": %u,\n", cpu->cred->uid.val);
        seq_printf(archivo, "\"state\":%d,\n", cpu->__state);
        mm = get_task_mm(cpu);
        if (mm)
        {
            seq_printf(archivo, "\"memory\":%lu,\n", get_mm_rss(mm));
            mmput(mm);
        }
        else
        {
            seq_printf(archivo, "\"memory\":%d,\n", 0);
        }
        seq_printf(archivo, "\"children\":[");
        list_for_each(lstProcess, &(cpu->children))
        {
            seq_printf(archivo, "\n{\n");
            child = list_entry(lstProcess, struct task_struct, sibling);
            seq_printf(archivo, "\"pid\":%d,\n", child->pid);
            seq_printf(archivo, "\"name\":\"%s\",\n", child->comm);
            seq_printf(archivo, "\"user\": %u,\n", child->cred->uid.val);
            seq_printf(archivo, "\"state\":%d,\n", child->__state);
            mm = get_task_mm(child);
            if (mm)
            {
                seq_printf(archivo, "\"memory\":%lu\n", get_mm_rss(mm));
                mmput(mm);
            }
            else
            {
                seq_printf(archivo, "\"memory\":%d\n", 0);
            }
            seq_printf(archivo, "}");
            if (lstProcess->next != &(cpu->children))
            {
                seq_printf(archivo, ",");
            }
        }
        seq_printf(archivo, "]\n},\n");
    }
    seq_printf(archivo, "{}\n]\n");
    return 0;
}

// Funcion que se ejecutara cada vez que se lea el archivo con el comando CAT
static int al_abrir(struct inode *inode, struct file *file)
{
    return single_open(file, escribir_archivo, NULL);
}

// Si el kernel es 5.6 o mayor se usa la estructura proc_ops
static struct proc_ops operaciones =
    {
        .proc_open = al_abrir,
        .proc_read = seq_read};

// Funcion a ejecuta al insertar el modulo en el kernel con insmod
static int _insert(void)
{
    proc_create("cpu_201709450", 0, NULL, &operaciones);
    printk(KERN_INFO "Marvin Daniel Rodriguez Felix\n");
    return 0;
}

// Funcion a ejecuta al remover el modulo del kernel con rmmod
static void _remove(void)
{
    remove_proc_entry("cpu_201709450", NULL);
    printk(KERN_INFO "Segundo Semestre 2022\n");
}

module_init(_insert);
module_exit(_remove);