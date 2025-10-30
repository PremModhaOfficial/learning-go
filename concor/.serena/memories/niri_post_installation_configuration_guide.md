# Niri Post-Installation Configuration and Customization Guide

## Research Summary

Comprehensive research on Niri desktop configuration covering all post-installation aspects:

## Key Configuration Details

**Primary Configuration:**
- Format: KDL (KDL Document Language)
- Location: ~/.config/niri/config.kdl  
- Addon configs: TOML format in ~/.config/niri/

**Core Configuration Elements:**
1. **Key Bindings**: KDL syntax with powerful customization
   ```kdl
   binds {
       workspace_previous "Mod4+j"
       workspace_next "Mod4+k"
       "Mod4+Return" spawn "alacritty"
   }
   ```

2. **Workspace Management**: Scrollable-tiling system
   - Infinite workspace within scrollable stacks
   - Automatic window sizing and positioning
   - Integration with niri-autoname-workspaces

3. **Window Rules**: Application-specific behavior
   ```kdl
   windows {
       class="firefox" floating=false
       class="zenity" floating=true
   }
   ```

4. **External Tool Integration**:
   - spawn-at-startup for daemons
   - IPC communication for tool coordination
   - Key bindings for interactive tools

## Essential Ecosystem Tools

**Core Tools:**
- **niri-autoname-workspaces**: Automatic workspace naming with icons
- **niriswitcher**: Application switcher with workspace awareness  
- **Waybar**: Status bar with Niri-specific modules
- **bemenu/tofi/rofi**: Application launchers
- **kanshi**: Multi-monitor display management

**Configuration Locations:**
- Niri: ~/.config/niri/config.kdl
- Waybar: ~/.config/waybar/config + style.css
- niriswitcher: ~/.config/niriswitcher/config.toml
- Addons: ~/.config/niri/[tool-name].toml

## Customization Categories

1. **Appearance**: GTK/Qt themes, Waybar styling, fonts, colors
2. **Productivity**: Key bindings, workspace organization, window rules
3. **Accessibility**: Screen readers, keyboard navigation, high contrast
4. **Performance**: Startup optimization, resource management
5. **Multi-monitor**: Display management, per-monitor configuration

## Research Sources

- Niri GitHub repository documentation
- niri-autoname-workspaces tool documentation
- niriswitcher application switcher docs
- Waybar Niri integration documentation
- KDL format specification
- Community Reddit discussions and configurations
- User configurations and best practices

## Key Insights

- Niri uses KDL for human-friendly configuration
- Rich ecosystem with active community development
- Scrollable-tiling paradigm changes desktop interaction
- Extensive customization through external tools
- Strong integration with Wayland ecosystem

See complete guide at: /home/prem-modha/projects/Motadata/gopher/concor/niri_configuration_customization_guide.md