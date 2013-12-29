package widgets

var gladestr = `<?xml version="1.0" encoding="UTF-8"?>
<interface>
  <!-- interface-requires gtk+ 3.0 -->
  <object class="GtkListStore" id="itemsListStore">
    <columns>
      <!-- column-name Items -->
      <column type="gchararray"/>
      <!-- column-name Complete -->
      <column type="gboolean"/>
      <!-- column-name Style -->
      <column type="gint"/>
      <!-- column-name Sensitive -->
      <column type="gboolean"/>
    </columns>
  </object>
  <object class="GtkWindow" id="mainWindow">
    <property name="can_focus">False</property>
    <property name="title" translatable="yes">Agenda</property>
    <property name="window_position">center</property>
    <property name="default_width">350</property>
    <property name="default_height">430</property>
    <child>
      <object class="GtkGrid" id="grid1">
        <property name="visible">True</property>
        <property name="can_focus">False</property>
        <property name="margin_left">12</property>
        <property name="margin_right">12</property>
        <property name="margin_top">12</property>
        <property name="margin_bottom">12</property>
        <property name="orientation">vertical</property>
        <property name="row_spacing">12</property>
        <child>
          <object class="GtkScrolledWindow" id="scrolledwindow1">
            <property name="visible">True</property>
            <property name="can_focus">False</property>
            <property name="vexpand">True</property>
            <property name="shadow_type">in</property>
            <child>
              <object class="GtkViewport" id="viewport1">
                <property name="visible">True</property>
                <property name="can_focus">False</property>
                <child>
                  <object class="GtkTreeView" id="itemsTreeView">
                    <property name="visible">True</property>
                    <property name="can_focus">False</property>
                    <property name="model">itemsListStore</property>
                    <property name="headers_visible">False</property>
                    <property name="headers_clickable">False</property>
                    <property name="reorderable">True</property>
                    <property name="search_column">1</property>
                    <child internal-child="selection">
                      <object class="GtkTreeSelection" id="treeview-selection"/>
                    </child>
                    <child>
                      <object class="GtkTreeViewColumn" id="completeColumn">
                        <property name="sizing">autosize</property>
                        <property name="title" translatable="yes">Complete</property>
                        <property name="sort_column_id">1</property>
                        <child>
                          <object class="GtkCellRendererToggle" id="completeToggleRenderer"/>
                          <attributes>
                            <attribute name="active">1</attribute>
                          </attributes>
                        </child>
                      </object>
                    </child>
                    <child>
                      <object class="GtkTreeViewColumn" id="itemsColumn">
                        <property name="title" translatable="yes">Items</property>
                        <property name="sort_column_id">1</property>
                        <child>
                          <object class="GtkCellRendererText" id="cellrenderertext1"/>
                          <attributes>
                            <attribute name="sensitive">3</attribute>
                            <attribute name="strikethrough">1</attribute>
                            <attribute name="style">2</attribute>
                            <attribute name="text">0</attribute>
                          </attributes>
                        </child>
                      </object>
                    </child>
                  </object>
                </child>
              </object>
            </child>
          </object>
          <packing>
            <property name="left_attach">0</property>
            <property name="top_attach">0</property>
            <property name="width">1</property>
            <property name="height">1</property>
          </packing>
        </child>
        <child>
          <object class="GtkEntry" id="itemTextEntry">
            <property name="visible">True</property>
            <property name="can_focus">True</property>
            <property name="hexpand">True</property>
            <property name="invisible_char">•</property>
            <property name="invisible_char_set">True</property>
            <property name="secondary_icon_stock">gtk-add</property>
            <property name="secondary_icon_tooltip_markup" translatable="yes">Add task</property>
            <property name="placeholder_text">Add a task</property>
          </object>
          <packing>
            <property name="left_attach">0</property>
            <property name="top_attach">1</property>
            <property name="width">1</property>
            <property name="height">1</property>
          </packing>
        </child>
      </object>
    </child>
  </object>
</interface>
`